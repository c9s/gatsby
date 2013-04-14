package sqlutils
import "fmt"

import "reflect"
import "strings"
import "strconv"
import "database/sql"

func BuildInsertColumnsFromMap(cols map[string]interface{}) (string, []interface{}) {
	var columnNames []string
	var valueFields []string
	var values		[]interface{}
	var i int = 1

	for col, arg := range cols {
		columnNames = append(columnNames, col)
		valueFields = append(valueFields, fmt.Sprintf("$%d", i))
		values    = append(values, arg)
		i++
	}
	sql := "( " + strings.Join(columnNames, ", ") + " )" +
		" VALUES ( " + strings.Join(valueFields, ", ") + " )"
	return sql , values
}


// Build Insert Column Clause from a struct type.
func BuildInsertColumns(val interface{}) (string, []interface{}) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var columnNames []string
	var valueFields []string
	var values      []interface{}
	var fieldId int = 1

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}

		var attributes = GetColumnAttributesFromTag(&tag)

		// if it's a serial column (with auto-increment, we can simply skip)
		if _, ok := attributes["serial"] ; ok {
			continue
		}

		columnNames = append(columnNames, *columnName)
		valueFields = append(valueFields, "$" + strconv.Itoa(fieldId) )
		values      = append(values, field.Interface() )
		fieldId++
	}
	return "( " + strings.Join(columnNames,", ") + " ) " +
	"VALUES ( " + strings.Join(valueFields,", ") + " )", values
}


func BuildInsertClause(val interface{}) (string, []interface{}) {
	tableName := GetTableName(val)
	sql, values := BuildInsertColumns(val)
	return "INSERT INTO " + tableName + sql, values
}

func GetReturningIdFromRows(rows * sql.Rows) (int64, error) {
	var id int64
	var err error
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}

