/*
 * Shift
 *
 * Copyright (c) 2017 QingCloud.com.
 * All rights reserved.
 *
 */

package shift

import (
	"fmt"
	"strings"
	"time"
	"xlog"

	"github.com/juju/errors"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
)

type Shift struct {
	log           *xlog.Log
	cfg           *Config
	toPool        *Pool
	fromPool      *Pool
	canal         *canal.Canal
	behindsTicker *time.Ticker
	done          chan bool
	handler       *EventHandler
	allDone       bool
	panicHandler  func(log *xlog.Log, format string, v ...interface{})
}

func NewShift(log *xlog.Log, cfg *Config) *Shift {
	log.Info("shift.cfg:%#v", cfg)
	return &Shift{
		log:           log,
		cfg:           cfg,
		done:          make(chan bool),
		behindsTicker: time.NewTicker(time.Duration(5000) * time.Millisecond),
		panicHandler:  logPanicHandler,
	}
}

func (shift *Shift) prepareConnection() error {
	log := shift.log
	cfg := shift.cfg

	fromPool, err := NewPool(log, 16, cfg.From, cfg.FromUser, cfg.FromPassword)
	if err != nil {
		log.Error("shift.start.from.connection.pool.error:%+v", err)
		return err
	}
	shift.fromPool = fromPool
	log.Info("shift.[%s].connection.done...", cfg.From)

	toPool, err := NewPool(log, cfg.Threads, cfg.To, cfg.ToUser, cfg.ToPassword)
	if err != nil {
		log.Error("shift.start.to.connection.pool.error:%+v", err)
		return err
	}
	shift.toPool = toPool
	log.Info("shift.[%s].connection.done...", cfg.To)
	log.Info("shift.prepare.connections.done...")
	return nil
}

func (shift *Shift) prepareTable() error {
	log := shift.log
	cfg := shift.cfg

	// From connection.
	fromConn := shift.fromPool.Get()
	defer shift.fromPool.Put(fromConn)

	// To connection.
	toConn := shift.toPool.Get()
	defer shift.toPool.Put(toConn)

	// Get databases
	log.Info("shift.get.database...")
	sql := "show databases;"
	r, err := fromConn.Execute(sql)
	if err != nil {
		log.Error("shift.check.database.sql[%s].error:%+v", sql, err)
		return err
	}
	for i := 0; i < r.RowNumber(); i++ {
		str, _ := r.GetString(i, 0)
		if _, isSystem := sysDatabases[strings.ToLower(str)]; !isSystem {
			cfg.Databases = append(cfg.Databases, str)
		}
	}
	if len(cfg.Databases) == 0 {
		return errors.New("no.database.to.shift")
	}

	for j := 0; j < len(cfg.Databases); j++ {
		// Prepare database, check the database is not system database and create them.
		log.Info("shift.prepare.database[%s]...", cfg.Databases[j])
		sql := fmt.Sprintf("select * from information_schema.tables where table_schema = '%s' limit 1", cfg.Databases[j])
		r, err := toConn.Execute(sql)
		if err != nil {
			log.Error("shift.check.database.sql[%s].error:%+v", sql, err)
			return err
		}

		if r.RowNumber() == 0 {
			sql := fmt.Sprintf("create database if not exists `%s`", cfg.Databases[j])
			if _, err := toConn.Execute(sql); err != nil {
				log.Error("shift.create.database.sql[%s].error:%+v", sql, err)
				return err
			}
			log.Info("shift.prepare.database.done...")
		} else {
			log.Info("shift.database.exists...")
		}

		// Get tables
		sql = fmt.Sprintf("use `%s`", cfg.Databases[j])
		r, err = fromConn.Execute(sql)
		if err != nil {
			log.Error("shift.check.database.sql[%s].error:%+v", sql, err)
			return err
		}

		sql = fmt.Sprintf("show tables")
		r, err = fromConn.Execute(sql)
		if err != nil {
			log.Error("shift.check.database.sql[%s].error:%+v", sql, err)
			return err
		}

		var tables []string
		for i := 0; i < r.RowNumber(); i++ {
			str, _ := r.GetString(i, 0)
			tables = append(tables, str)
		}
		if len(tables) == 0 {
			log.Error("shift.check.database.[%+v].no.tables", cfg.Databases[j])
			continue // don`t need return err
		}

		// prepare tables
		for i := 0; i < len(tables); i++ {
			log.Info("shift.prepare.table[%s/%s]...", cfg.Databases[j], tables[i])
			sql = fmt.Sprintf("show create table `%s`.`%s`", cfg.Databases[j], tables[i])
			r, err = fromConn.Execute(sql)
			if err != nil {
				log.Error("shift.show.[%s].create.table.sql[%s].error:%+v", cfg.From, sql, err)
				return err
			}
			sql, err = r.GetString(0, 1)
			if err != nil {
				log.Error("shift.show.[%s].create.table.get.error:%+v", cfg.From, err)
				return err
			}
			sql = strings.Replace(sql, fmt.Sprintf("CREATE TABLE `%s`", tables[i]), fmt.Sprintf("CREATE TABLE `%s`.`%s`", cfg.Databases[j], tables[i]), 1)
			if _, err := toConn.Execute(sql); err != nil {
				log.Error("shift.create.[%s].table.sql[%s].error:%+v", cfg.From, sql, err)
				return err
			}
			log.Info("shift.prepare.table.done...")
		}
	}
	return nil
}

func (shift *Shift) prepareCanal() error {
	log := shift.log
	conf := shift.cfg
	cfg := canal.NewDefaultConfig()
	cfg.Addr = conf.From
	cfg.User = conf.FromUser
	cfg.Password = conf.FromPassword
	cfg.Dump.ExecutionPath = conf.MySQLDump
	cfg.Dump.DiscardErr = false
	// TableDB and Tables will be used in the future
	// cfg.Dump.TableDB = conf.FromDatabase
	// cfg.Dump.Tables = []string{conf.FromTable}
	cfg.Dump.Databases = conf.Databases
	log.Info("gry+++数组:%+v", cfg.Dump.Databases)

	// canal
	canal, err := canal.NewCanal(cfg)
	if err != nil {
		log.Error("shift.canal.new.error:%+v", err)
		return err
	}

	handler := NewEventHandler(log, shift)
	canal.SetEventHandler(handler)
	shift.handler = handler
	shift.canal = canal
	if err := canal.Start(); err != nil {
		log.Error("shift.canal.start.error:%+v", err)
		return err
	}
	log.Info("shift.prepare.canal.done...")
	return nil
}

/*
	mysql> checksum table sbtest.sbtest1;
	+----------------+-----------+
	| Table          | Checksum  |
	+----------------+-----------+
	| sbtest.sbtest1 | 410139351 |
	+----------------+-----------+
*/
// ChecksumTable ensure that FromTable and ToTable are consistent
func (shift *Shift) ChecksumTable() error {
	log := shift.log
	var fromchecksum, tochecksum uint64

	if _, isSystem := sysDatabases[strings.ToLower(shift.cfg.FromDatabase)]; isSystem {
		log.Info("shift.checksum.table.skip.system.table[%s.%s]", shift.cfg.FromDatabase, shift.cfg.FromTable)
		return nil
	}

	checksumFunc := func(t string, Conn *client.Conn, Database string, Table string, c chan uint64) {
		sql := fmt.Sprintf("checksum table %s.%s", Database, Table)
		r, err := Conn.Execute(sql)
		if err != nil {
			shift.panicMe("shift.checksum.%s.table[%s.%s].error:%+v", t, Database, Table, err)
		}

		v, err := r.GetUint(0, 1)
		if err != nil {
			shift.panicMe("shift.get.%s.table[%s.%s].checksum.error:%+v", Database, Table, err)
		}
		c <- v
	}

	fromchan := make(chan uint64, 1)
	tochan := make(chan uint64, 1)

	// execute checksum func
	{
		fromConn := shift.fromPool.Get()
		defer shift.fromPool.Put(fromConn)

		toConn := shift.toPool.Get()
		defer shift.toPool.Put(toConn)

		go checksumFunc("from", fromConn, shift.cfg.FromDatabase, shift.cfg.FromTable, fromchan)
		go checksumFunc("to", toConn, shift.cfg.ToDatabase, shift.cfg.ToTable, tochan)
	}
	fromchecksum = <-fromchan
	tochecksum = <-tochan

	if fromchecksum != tochecksum {
		err := fmt.Errorf("checksum not equivalent: from-table checksum is %v, to-table checksum is %v", fromchecksum, tochecksum)
		log.Error("shift.checksum.table.err:%+v", err)
		return err
	}
	return nil
}

/*
   mysql> show master status;
   +------------------+-----------+--------------+------------------+------------------------------------------------+
   | File             | Position  | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set                              |
   +------------------+-----------+--------------+------------------+------------------------------------------------+
   | mysql-bin.000002 | 112107994 |              |                  | 4dc59763-5431-11e7-90cb-5254281e57de:1-2561361 |
   +------------------+-----------+--------------+------------------+------------------------------------------------+
*/
func (shift *Shift) masterPosition() *mysql.Position {
	position := &mysql.Position{}

	fromConn := shift.fromPool.Get()
	defer shift.fromPool.Put(fromConn)

	sql := "show master status"
	r, err := fromConn.Execute(sql)
	if err != nil {
		shift.panicMe("shift.get.master[%s].postion.error:%+v", shift.cfg.From, err)
		return position
	}

	file, err := r.GetString(0, 0)
	if err != nil {
		shift.panicMe("shift.get.master[%s].file.error:%+v", shift.cfg.From, err)
		return position
	}

	pos, err := r.GetUint(0, 1)
	if err != nil {
		shift.panicMe("shift.get.master[%s].pos.error:%+v", shift.cfg.From, err)
		return position
	}
	position.Name = file
	position.Pos = uint32(pos)
	return position
}

func (shift *Shift) behindsCheckStart() error {
	go func(s *Shift) {
		log := s.log
		log.Info("shift.dumping...")
		<-s.canal.WaitDumpDone()
		// Wait dump worker done.
		log.Info("shift.wait.dumper.background.worker...")
		shift.handler.WaitWorkerDone()
		log.Info("shift.wait.dumper.background.worker.done...")
		log.Info("shift.set.dumper.background.worker.done...")
		s.canal.SetDumpWorkerDone()
		log.Info("shift.dumping.done...")

		for range s.behindsTicker.C {
			masterPos := s.masterPosition()
			syncPos := s.canal.SyncedPosition()
			behinds := int(masterPos.Pos - syncPos.Pos)
			log.Info("--shift.check.behinds[%d]--master[%+v]--synced[%+v]", behinds, masterPos, syncPos)
			if behinds <= shift.cfg.Behinds {
				shift.setRadon()
				return
			} else {
				factor := float32(shift.cfg.Behinds+1) / float32(behinds+1)
				log.Info("shift.set.throttle.behinds[%v].cfgbehinds[%v].factor[%v]", behinds, shift.cfg.Behinds, factor)
				shift.setRadonThrottle(factor)
			}
		}
	}(shift)
	return nil
}

// Start used to start canal and behinds ticker.
func (shift *Shift) Start() error {
	if err := shift.prepareConnection(); err != nil {
		return err
	}
	if err := shift.prepareTable(); err != nil {
		return err
	}
	if err := shift.prepareCanal(); err != nil {
		return err
	}
	if err := shift.behindsCheckStart(); err != nil {
		return err
	}
	return nil
}

// Close used to destroy all the resource.
func (shift *Shift) Close() {
	shift.behindsTicker.Stop()
	shift.fromPool.Close()
	shift.toPool.Close()
	shift.canal.Close()
	shift.Cleanup()
}

func (shift *Shift) Done() chan bool {
	return shift.done
}

func (shift *Shift) panicMe(format string, v ...interface{}) {
	shift.Cleanup()
	shift.panicHandler(shift.log, format, v)
}
