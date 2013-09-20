package sqlutils

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
		valueFields = append(valueFields, "$"+strconv.Itoa(i))
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

	var columnNamesSql string = ""
	var valueFieldsSql string = ""

	var values []interface{}
	var fieldId int = 1

	for i := 0; i < t.NumField(); i++ {
		var fieldType = typeOfT.Field(i)
		var tag reflect.StructTag = fieldType.Tag

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}

		if HasColumnAttributeFromTag(&tag, "serial") {
			continue
		}

		var field reflect.Value = t.Field(i)
		var val interface{} = field.Interface()

		// if time is null or with zero value, just skip it.
		switch tv := val.(type) {
		case *time.Time:
			if tv == nil || tv.Unix() == -62135596800 {
				continue
			}
		case time.Time:
			if tv.Unix() == -62135596800 {
				continue
			}
		case string:
			if HasColumnAttributeFromTag(&tag, "date") && tv == "" {
				continue
			}
		}

		columnNamesSql += *columnName + ", "
		valueFieldsSql += "$" + strconv.Itoa(fieldId) + ", "
		values = append(values, val)
		fieldId++
	}
	return "( " + columnNamesSql[:len(columnNamesSql)-2] + " ) " +
		"VALUES ( " + valueFieldsSql[:len(valueFieldsSql)-2] + " )", values
}

func BuildInsertClause(val interface{}) (string, []interface{}) {
	sql, values := BuildInsertColumns(val)
	return "INSERT INTO " + GetTableName(val) + sql, values
}
