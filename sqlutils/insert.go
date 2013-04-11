package sqlutils
// import "fmt"

import "reflect"
import "strings"
import "strconv"
import "github.com/c9s/inflect"

func BuildInsertColumnClause(val interface{}) (string, []interface{}) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()
	tableName := inflect.Tableize(typeOfT.Name())

	var columnNames []string
	var valueFields []string
	var values      []interface{}

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)
		// var fieldType  reflect.Type = field.Type()

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		// fieldAttrs := GetColumnAttributesFromTag(&tag)
		columnNames = append(columnNames, *columnName)
		valueFields = append(valueFields, "$" + strconv.Itoa(i + 1) )
		values      = append(values, field.Interface() )
	}
	return "INSERT INTO " + tableName + " (" + strings.Join(columnNames,",") + ") " +
		" VALUES (" + strings.Join(valueFields,",") + ")", values
}

