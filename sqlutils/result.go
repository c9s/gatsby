package sqlutils

type Result struct {
	Sql string
	Id  int
	Error error
}

func NewErrorResult(err error) (*Result) {
	r := Result{Error: err}
	return &r
}

func NewResult(sql string) (*Result) {
	r := Result{Sql: sql}
	return &r
}

