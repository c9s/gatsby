package sqlutils
import "database/sql"

type Result struct {
	Sql string
	Id  int64
	Error error
	Result sql.Result
}

// Return Error Result, which is used in Create, Update, Delete functions.
func NewErrorResult(err error,sql string) (*Result) {
	r := Result{Error: err, Sql: sql}
	return &r
}

// Create new result object with SQL statement string.
func NewResult(sql string) (*Result) {
	r := Result{Sql: sql}
	return &r
}

