package sqlutils
import "fmt"
import "reflect"
import "strings"
import "database/sql"


// This function generates "UPDATE {table} SET name = $1, name2 = $2"
func BuildUpdateClause(val interface{}) (string, []interface{}) {
	tableName := GetTableName(val)
	sql, values := BuildUpdateColumns(val)
	return "UPDATE " + tableName + " SET " + sql, values
}


// This function builds update columns from a map
// which generates SQL like "name = $1, phone = $2".
func BuildUpdateColumnsFromMap(cols map[string]interface{}) (string, []interface{}) {
	var setFields []string
	var values      []interface{}
	var i int = 1
	for col, arg := range cols {
		setFields = append(setFields, fmt.Sprintf("%s = $%d", col, i) )
		values    = append(values, arg)
		i++
	}
	return strings.Join(setFields, ", "), values
}


// This function generate update columns from a struct object.
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



func Update(db *sql.DB, val interface{}) (*Result) {

	pkName := GetPrimaryKeyColumnName(&val)
	if pkName == nil {
		panic("primary key column is not defined.")
	}

	sql, values := BuildUpdateClause(val)

	if val.(PrimaryKey) != nil {
		id := val.(PrimaryKey).GetPkId()
		values = append(values, id)
	}

	sql += fmt.Sprintf(" WHERE %s = $%d", *pkName, len(values))

	stmt, err := db.Prepare(sql)

	defer func() { stmt.Close() }()

	if err != nil {
		return NewErrorResult(err,sql)
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	result := NewResult(sql)
	result.Result = res
	return result
}

