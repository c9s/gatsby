package sqlutils
// import "fmt"

import "reflect"
import "strings"
import "strconv"
import "github.com/c9s/inflect"

func BuildInsertColumnClause(val interface{}) (string) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()
	tableName := inflect.Tableize(typeOfT.Name())
	columns := ParseColumnNames(val)
	var valueFields []string
	for i, _ := range columns {
		valueFields = append(valueFields, "$" + strconv.Itoa(i + 1) )
	}
	return "INSERT INTO " + tableName + " (" + strings.Join(columns,",") + ") " +
		" VALUES (" + strings.Join(valueFields,",") + ")"
}

