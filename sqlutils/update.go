package sqlutils
// import "fmt"

import "reflect"
import "strings"
import "strconv"
import "github.com/c9s/inflect"


// Generate "UPDATE {table} SET name = $1, name2 = $2"
func BuildUpdateColumnClause(val interface{}) (string) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()
	tableName := inflect.Tableize(typeOfT.Name())
	return "UPDATE " + tableName + " SET " + BuildUpdateColumns(val)
}

func BuildUpdateColumns(val interface{}) (string) {
	columns := ParseColumnNames(val)
	var fields []string
	for i, name := range columns {
		fields = append(fields, name + " = $" + strconv.Itoa(i + 1) )
	}
	return strings.Join(fields,", ")
}


