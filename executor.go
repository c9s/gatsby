package gatsby

import "database/sql"

type Executor interface {
	Exec(string, ...interface{}) (*sql.Result, error)
}
