package sqlutils

type Result struct {
	Sql string
	Id  int
	Error error
}

func NewErrorResult(err error,sql string) (*Result) {
	r := Result{Error: err, Sql: sql}
	return &r
}

func NewResult(sql string) (*Result) {
	r := Result{Sql: sql}
	return &r
}

