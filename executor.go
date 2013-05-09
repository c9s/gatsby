package gatsby

import "database/sql"

type Executor interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(string) (*sql.Stmt, error)
	QueryRow(string, ...interface{}) *sql.Row
}
