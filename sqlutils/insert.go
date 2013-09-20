package sqlutils

import "fmt"
import "reflect"
import "strings"
import "strconv"
import "time"

func BuildInsertColumnsFromMap(cols map[string]interface{}) (string, []interface{}) {
	var columnNames []string
	var valueFields []string
	var values []interface{}
	var i int = 1

	for col, arg := range cols {
		columnNames = append(columnNames, col)
		valueFields = append(valueFields, fmt.Sprintf("$%d", i))
		values = append(values, arg)
		i++
	}
	return "( " + strings.Join(columnNames, ", ") + " )" +
		" VALUES ( " + strings.Join(valueFields, ", ") + " )", values
}

// Build Insert Column Clause from a struct type.
func BuildInsertColumns(val interface{}) (string, []interface{}) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var columnNames []string
	var valueFields []string
	var values []interface{}
	var fieldId int = 1

	for i := 0; i < t.NumField(); i++ {
		var fieldType = typeOfT.Field(i)
		var tag reflect.StructTag = fieldType.Tag

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}

		var attributes = GetColumnAttributesFromTag(&tag)

		// if it's a serial column (with auto-increment, we can simply skip)
		if _, ok := attributes["serial"]; ok {
			continue
		}

		var field reflect.Value = t.Field(i)
		var val interface{} = field.Interface()

		// if time is null or with zero value, just skip it.
		if fieldType.Type.String() == "*time.Time" {
			if timeVal, ok := val.(*time.Time); ok {
				if timeVal == nil || timeVal.Unix() == -62135596800 {
					continue
				}
			}
		}

		if attributes["date"] {
			switch val.(type) {
			case string:
				if val == "" {
					continue
				}
			}
		}

		columnNames = append(columnNames, *columnName)
		valueFields = append(valueFields, "$"+strconv.Itoa(fieldId))
		values = append(values, val)
		fieldId++
	}
	return "( " + strings.Join(columnNames, ", ") + " ) " +
		"VALUES ( " + strings.Join(valueFields, ", ") + " )", values
}

func BuildInsertClause(val interface{}) (string, []interface{}) {
	sql, values := BuildInsertColumns(val)
	return "INSERT INTO " + GetTableName(val) + sql, values
}
