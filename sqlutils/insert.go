package sqlutils
// import "fmt"
// import "reflect"
// import "strings"

func BuildInsertColumnClause(val interface{}) {
	columns := ParseColumnNames(val)
	_ = columns
}

