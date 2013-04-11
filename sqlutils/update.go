package sqlutils
import "fmt"
import "reflect"
import "strings"
import "database/sql"
import "errors"


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



func Update(db *sql.DB, val interface{}) (*Result) {

	if val.(PrimaryKey) == nil {
		return NewErrorResult(errors.New("PrimaryKey interface is required."),"")
	}

	sql, values := BuildUpdateClause(val)

	id := val.(PrimaryKey).GetPkId()
	values = append(values, id)

	sql += fmt.Sprintf(" WHERE id = $%d", len(values))

	stmt, err := db.Prepare(sql)
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


