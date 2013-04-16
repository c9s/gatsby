package sqlutils
import "strings"
import "database/sql"
import "reflect"
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
	value := reflect.Indirect( reflect.ValueOf(val) )
	typeOfVal := value.Type()
	sliceOfVal := reflect.SliceOf(typeOfVal)
	var slice = reflect.MakeSlice(sliceOfVal,0,100)
	defer func() { rows.Close() }()
	for rows.Next() {
		var newValue = reflect.New(typeOfVal)
		err = FillFromRow(newValue.Interface() , rows)
		if err != nil {
			return slice.Interface(), NewErrorResult(err, sql)
		}
		slice = reflect.Append(slice, reflect.Indirect(newValue) )
	}
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

