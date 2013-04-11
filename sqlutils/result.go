package sqlutils
import "database/sql"

type Result struct {
	Sql string
	Id  int
	Error error
	Result sql.Result
}

func NewErrorResult(err error,sql string) (*Result) {
	r := Result{Error: err, Sql: sql}
	return &r
}

func NewResult(sql string) (*Result) {
	r := Result{Sql: sql}
	return &r
}

