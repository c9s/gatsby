package sqlutils
import "strings"
import "database/sql"
import "reflect"
// import "fmt"
// import "errors"

// Generate SQL columns string for selecting.
func BuildSelectColumnClause(val interface{}) (string) {
	columns := ParseColumnNames(val)
	return strings.Join(columns,",")
}

func BuildSelectClause(val interface{}) (string) {
	// get table name
	// inflect.Underscore()
	tableName := GetTableName(val)
	return "SELECT " + BuildSelectColumnClause(val) + " FROM " + tableName;
}


func SelectQuery(db *sql.DB, val interface{}) (*sql.Rows, error) {
	sql := BuildSelectClause(val)
	return PrepareAndQuery(db, sql)
}

func SelectQueryWith(db *sql.DB, val interface{}, postSql string, args ...interface{}) (*sql.Rows, error) {
	sql := BuildSelectClause(val) + " " + postSql
	return PrepareAndQuery(db, sql, args)
}

func Select(db *sql.DB, val interface{}) (interface{}, *Result) {
	sql := BuildSelectClause(val)
	rows, err := PrepareAndQuery(db, sql)

	if err != nil {
		return nil, NewErrorResult(err,sql)
	}

	value := reflect.ValueOf(val)
	typeOfVal := value.Type()

	sliceOfVal := reflect.SliceOf(typeOfVal)
	var slice = reflect.MakeSlice(sliceOfVal,0,100)
	// var slice []interface{}

	defer func() { rows.Close() }()

	for rows.Next() {
		var newVal = reflect.Indirect( reflect.New(typeOfVal) )

		/*
		var val = newVal.Elem().Interface()
		fmt.Println( val )
		*/


		// FillFromRow(newVal.Elem().Interface() ,rows)
		/*
		if err != nil {
			return items.Interface().(*[]interface{}), NewErrorResult(err,sql)
		}
		*/
		// slice = reflect.Append(slice, newVal)
		_ = newVal
	}
	// return slice.Interface().(*[]interface{}), NewResult(sql)
	// return slice.Slice(0,30), NewResult(sql)
	return slice.Interface(), NewResult(sql)
}


func SelectWith(db *sql.DB, val interface{}, postSql string, args ...interface{}) (*[]interface{}, *Result) {
	sql := BuildSelectClause(val) + " " + postSql
	rows, err := PrepareAndQuery(db, sql, args)

	if err != nil {
		return nil, NewErrorResult(err,sql)
	}

	var items = new([]interface{})

	defer func() { rows.Close() }()
	return items, NewResult(sql)
	/*
	if rows.Next() {
		err = FillFromRow(val,rows)
		if err != nil {
			return items, NewErrorResult(err,sql)
		}
	}
	*/
	return items, NewResult(sql)
}

