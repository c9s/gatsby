package sqlutils

import "strings"
import "reflect"
import "database/sql"

// import "errors"

// Build SQL columns string for selecting,
// this function returns "column1, column2, column3"
func BuildSelectColumnClauseFromStruct(val interface{}) string {
	var columns = ReflectColumnNames(val)
	return strings.Join(columns, ", ")
}

// Build SQL columns string for selecting, this function returns
// "alias.column1, alias.column2, alias.column3"
func BuildSelectColumnClauseFromStructWithAlias(val interface{}, alias string) string {
	var columns = ReflectColumnNames(val)
	var aliasColumns = []string{}
	for _, n := range columns {
		aliasColumns = append(aliasColumns, alias+"."+n)
	}
	return strings.Join(aliasColumns, ", ")
}

// Given a struct object, return a "SELECT ... FROM {tableName}" SQL clause.
func BuildSelectClause(val interface{}) string {
	tableName := GetTableName(val)
	return "SELECT " + BuildSelectColumnClauseFromStruct(val) + " FROM " + tableName
}

// Given a struct object and an alias string,
// This function returns a "SELECT alias1.column1, alias1.column2 FROM tableName alias" clause
func BuildSelectClauseWithAlias(val interface{}, alias string) string {
	tableName := GetTableName(val)
	return "SELECT " + BuildSelectColumnClauseFromStructWithAlias(val, alias) + " FROM " + tableName + " " + alias
}

func CreateStructSliceFromRows(val interface{}, rows *sql.Rows) (interface{}, error) {
	value := reflect.Indirect(reflect.ValueOf(val))
	typeOfVal := value.Type()
	sliceOfVal := reflect.SliceOf(typeOfVal)
	var slice = reflect.MakeSlice(sliceOfVal, 0, 200)
	defer rows.Close()
	for rows.Next() {
		var newValue = reflect.New(typeOfVal)
		var err = FillFromRow(newValue.Interface(), rows)
		if err != nil {
			return slice.Interface(), err
		}
		slice = reflect.Append(slice, reflect.Indirect(newValue))
	}
	err := rows.Err()
	if err != nil {
		return slice, err
	}
	return slice.Interface(), nil
}
