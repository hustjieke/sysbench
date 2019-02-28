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

	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/client"
)

func (h *EventHandler) DeleteRow(e *canal.RowsEvent) {
	var conn *client.Conn
	log := h.log

	h.wg.Add(1)
	executeFunc := func(conn *client.Conn) {
		defer h.wg.Done()
		var keep = true

		pks := e.Table.PKColumns
		for i, row := range e.Rows {
			var values []string

			// keep connection in the loop, just put conn to pool when execute the last row
			if (i + 1) == len(e.Rows) {
				keep = false
			}

			// We have pk columns.
			if len(pks) > 0 {
				for _, pk := range pks {
					v := row[pk]
					if _, ok := v.([]byte); ok {
						values = append(values, fmt.Sprintf("%s=%q", e.Table.Columns[pk].Name, v))
					} else {
						values = append(values, fmt.Sprintf("%s=%#v", e.Table.Columns[pk].Name, v))
					}
				}
			} else {
				for j, v := range row {
					if v == nil {
						continue
					}
					if _, ok := v.([]byte); ok {
						values = append(values, fmt.Sprintf("%s=%q", e.Table.Columns[j].Name, v))
					} else {
						values = append(values, fmt.Sprintf("%s=%#v", e.Table.Columns[j].Name, v))
					}
				}
			}

			query := &Query{
				sql:       fmt.Sprintf("delete from `%s`.`%s` where %s", e.Table.Schema, e.Table.Name, strings.Join(values, " and ")),
				typ:       QueryType_DELETE,
				skipError: false,
			}
			log.Debug("----no:%d, query:%+v", i, query)
			h.execute(conn, keep, query)
		}
	}

	if h.xaConn != nil {
		conn = h.xaConn
	} else {
		conn = h.shift.toPool.Get()
	}

	executeFunc(conn)
}
