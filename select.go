package gatsby

import "database/sql"
import "gatsby/sqlutils"
import "reflect"

func CreateStructSliceFromRows(val interface{}, rows *sql.Rows) (interface{}, error) {
	value := reflect.Indirect(reflect.ValueOf(val))
	typeOfVal := value.Type()
	sliceOfVal := reflect.SliceOf(typeOfVal)
	var slice = reflect.MakeSlice(sliceOfVal, 0, 200)
	defer rows.Close()
	for rows.Next() {
		var newValue = reflect.New(typeOfVal)
		var err = FillFromRows(newValue.Interface(), rows)
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

/*
Select all records from a table which based on the record struct.
*/
func Select(db *sql.DB, val PtrRecord) (interface{}, *Result) {
	var sql = sqlutils.BuildSelectClause(val)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}
	defer rows.Close()

	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}

/*
Execute a select query to the database connection.
*/
func QuerySelect(db *sql.DB, val PtrRecord) (*sql.Rows, error) {
	sql := sqlutils.BuildSelectClause(val)
	return db.Query(sql)
}

func QuerySelectWith(db *sql.DB, val PtrRecord, postSql string, args ...interface{}) (*sql.Rows, error) {
	sql := sqlutils.BuildSelectClause(val) + " " + postSql
	return db.Query(sql, args...)
}

// Select a table and returns objects
func SelectWith(db *sql.DB, val PtrRecord, postSql string, args ...interface{}) (interface{}, *Result) {
	sql := sqlutils.BuildSelectClause(val) + " " + postSql
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}
	defer rows.Close()

	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}

func SelectWhere(db *sql.DB, val PtrRecord, conds WhereMap) (interface{}, *Result) {
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(conds)
	sql := sqlutils.BuildSelectClause(val) + whereSql
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}
	defer rows.Close()

	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}

func SelectFromQuery(db *sql.DB, val PtrRecord, sql string, args ...interface{}) (interface{}, *Result) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}
	defer rows.Close()

	slice, err := CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}
