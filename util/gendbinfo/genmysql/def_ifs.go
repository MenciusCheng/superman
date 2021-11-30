package genmysql

import "database/sql"

type MySqlDB interface {
	RawRows(sql string, values ...interface{}) (*sql.Rows, error)
	RawScan(dest interface{}, sql string, values ...interface{}) error
}
