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
	return db.Query(sql)
}

func SelectQueryWith(db *sql.DB, val interface{}, postSql string, args ...interface{}) (*sql.Rows, error) {
	sql := BuildSelectClause(val) + " " + postSql
	return db.Query(sql, args...)
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


func Select(db *sql.DB, val interface{}) (interface{}, *Result) {
	sql := BuildSelectClause(val)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, NewErrorResult(err,sql)
	}
	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err,sql)
	}
	return slice, NewResult(sql)
}


// select table with a postSQL
func SelectWith(db *sql.DB, val interface{}, postSql string, args ...interface{}) (interface{}, *Result) {
	sql := BuildSelectClause(val) + " " + postSql
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err,sql)
	}

	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err,sql)
	}
	return slice, NewResult(sql)
}

func SelectWhere(db *sql.DB, val interface{}, conds map[string]interface{}) (interface{}, *Result) {
	whereSql, args := BuildWhereClauseWithAndOp(conds)
	sql := BuildSelectClause(val) + whereSql
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err,sql)
	}

	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err,sql)
	}
	return slice, NewResult(sql)
}

func SelectFromQuery(db *sql.DB, val interface{}, sql string, args ...interface{} ) (interface{}, *Result) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err,sql)
	}
	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err,sql)
	}
	return slice, NewResult(sql)
}
