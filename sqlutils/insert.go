package sqlutils

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	NUMBER_HOLDER = iota
	QMARK_HOLDER
)

func BuildInsertColumnsFromMap(cols map[string]interface{}, placeholder int) (string, []interface{}) {
	var columnNames []string
	var valueFields []string
	var values []interface{}
	var i int = 1

	for col, arg := range cols {
		columnNames = append(columnNames, col)
		if placeholder == QMARK_HOLDER {
			valueFields = append(valueFields, "?")
		} else {
			valueFields = append(valueFields, "$"+strconv.Itoa(i))
		}
		values = append(values, arg)
		i++
	}
	return "( " + strings.Join(columnNames, ", ") + " )" +
		" VALUES ( " + strings.Join(valueFields, ", ") + " )", values
}

// Build Insert Column Clause from a struct type.
func BuildInsertColumns(val interface{}, placeholder int) (string, []interface{}) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var columnNamesSql string = ""
	var valueFieldsSql string = ""

	var values []interface{}
	var fieldId int = 1
	var tag reflect.StructTag
	var columnName *string
	var field reflect.Value
	var fieldType reflect.StructField

	for i := 0; i < t.NumField(); i++ {
		fieldType = typeOfT.Field(i)
		tag = fieldType.Tag

		if columnName = GetColumnNameFromTag(&tag); columnName == nil {
			continue
		}
		if HasColumnAttributeFromTag(&tag, "serial") {
			continue
		}

		field = t.Field(i)

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

		if placeholder == QMARK_HOLDER {
			valueFieldsSql += "?, "
		} else {
			valueFieldsSql += "$" + strconv.Itoa(fieldId) + ", "
		}
		values = append(values, val)
		fieldId++
	}
	return "( " + columnNamesSql[:len(columnNamesSql)-2] + " ) " +
		"VALUES ( " + valueFieldsSql[:len(valueFieldsSql)-2] + " )", values
}

func BuildInsertClause(val interface{}, placeholder int) (string, []interface{}) {
	sql, values := BuildInsertColumns(val, placeholder)
	return "INSERT INTO " + GetTableName(val) + sql, values
}
