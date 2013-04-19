package sqlutils
import "strings"
import "reflect"
import "database/sql"

// import "errors"

// Generate SQL columns string for selecting.
func BuildSelectColumnClauseFromStruct(val interface{}) (string) {
	columns := ReflectColumnNames(val)
	return strings.Join(columns,",")
}

func BuildSelectClause(val interface{}) (string) {
	// get table name
	// inflect.Underscore()
	tableName := GetTableName(val)
	return "SELECT " + BuildSelectColumnClauseFromStruct(val) + " FROM " + tableName;
}




func CreateStructSliceFromRows(val interface{}, rows *sql.Rows ) (interface{}, error) {
	value := reflect.Indirect( reflect.ValueOf(val) )
	typeOfVal := value.Type()
	sliceOfVal := reflect.SliceOf(typeOfVal)
	var slice = reflect.MakeSlice(sliceOfVal,0,200)
	defer func() { rows.Close() }()
	for rows.Next() {
		var newValue = reflect.New(typeOfVal)
		var err = FillFromRow(newValue.Interface() , rows)
		if err != nil {
			return slice.Interface(), err
		}
		slice = reflect.Append(slice, reflect.Indirect(newValue) )
	}
	err := rows.Err()
	if err != nil {
		return slice, err
	}
	return slice.Interface(), nil
}


