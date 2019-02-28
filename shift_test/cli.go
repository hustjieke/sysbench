/*
 * Shift
 *
 * Copyright (c) 2017 QingCloud.com.
 * All rights reserved.
 *
 */

package main

import (
	"build"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"shift"
	"syscall"
	"xlog"
)

var (
	from         = flag.String("from", "", "Source MySQL backend")
	fromUser     = flag.String("from-user", "", "MySQL user, must have replication privilege")
	fromPassword = flag.String("from-password", "", "MySQL user password")
	fromDatabase = flag.String("from-database", "", "Source database")
	fromTable    = flag.String("from-table", "", "Source table")

	to         = flag.String("to", "", "Destination MySQL backend")
	toUser     = flag.String("to-user", "", "MySQL user, must have replication privilege")
	toPassword = flag.String("to-password", "", "MySQL user password")
	toDatabase = flag.String("to-database", "", "Destination database")
	toTable    = flag.String("to-table", "", "Destination table")

	cleanup   = flag.Bool("cleanup", true, "Cleanup the from table after shifted(defaults true)")
	checksum  = flag.Bool("checksum", true, "Checksum the from table and to table after shifted(defaults true)")
	mysqlDump = flag.String("mysqldump", "mysqldump", "mysqldump path")
	threads   = flag.Int("threads", 16, "shift threads num(defaults 16)")
	behinds   = flag.Int("behinds", 2048, "seconds behinds num(default 2048)")
	radonURL  = flag.String("radonurl", "http://127.0.0.1:8080", "Radon RESTful api(defaults http://127.0.0.1:8080)")
)

func check(log *xlog.Log) {
	if *from == "" || *fromUser == "" || *to == "" || *toUser == "" {
		log.Panic("usage: shift --from=[host:port] --from-password=[password] -- to=[host:port] --to-user=[user] --to-password=[password] --cleanup=[false|true]")
	}
}

func main() {
	log := xlog.NewStdLog(xlog.Level(xlog.INFO))
	runtime.GOMAXPROCS(runtime.NumCPU())

	build := build.GetInfo()
	fmt.Printf("shift:[%+v]\n", build)

	// flags.
	flag.Parse()
	check(log)
	fmt.Println(`
           IMPORTANT: Please check that the shift run completes successfully.
           At the end of a successful shift run prints "shift.completed.OK!".`)

	cfg := &shift.Config{
		From:         *from,
		FromUser:     *fromUser,
		FromPassword: *fromPassword,
		FromDatabase: *fromDatabase,
		FromTable:    *fromTable,
		To:           *to,
		ToUser:       *toUser,
		ToPassword:   *toPassword,
		ToDatabase:   *toDatabase,
		ToTable:      *toTable,
		Cleanup:      *cleanup,
		MySQLDump:    *mysqlDump,
		Threads:      *threads,
		Behinds:      *behinds,
		RadonURL:     *radonURL,
		Checksum:     *checksum,
	}
	log.Info("shift.cfg:%+v", cfg)
	shift := shift.NewShift(log, cfg)
	if err := shift.Start(); err != nil {
		log.Panicf("shift.start.error:%+v", err)
	}
	defer shift.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	select {
	case <-shift.Done():
		fmt.Println("shift.completed.OK!")
	case <-sc:
		fmt.Println("shift.signal.done...")
	}
}
