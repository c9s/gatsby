package sqlutils
import "fmt"
import "reflect"
import "strings"


// Generate "UPDATE {table} SET name = $1, name2 = $2"
func BuildUpdateClause(val interface{}) (string, []interface{}) {
	tableName := GetTableName(val)
	sql, values := BuildUpdateColumns(val)
	return "UPDATE " + tableName + " SET " + sql, values
}

func BuildUpdateColumns(val interface{}) (string, []interface{}) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()
	var setFields []string
	var values      []interface{}

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		setFields = append(setFields, fmt.Sprintf("%s = $%d", *columnName, i + 1) )
		values    = append(values, field.Interface() )
	}
	return strings.Join(setFields,", "), values
}

// id, err := sqlutils.Create(struct pointer)
/*
func Update(db *sql.DB, val interface{}) (int,error) {
	sql , args := BuildUpdateClause(val)

	// for pgsql only
	sql += " RETURNING id"

	err := CheckRequired(val)
	if err != nil {
		return -1, err
	}

	rows, err := PrepareAndQuery(db,sql,args...)
	if err != nil {
		return -1, err
	}
	return GetReturningIdFromRows(rows)
}
*/


