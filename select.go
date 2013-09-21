package gatsby

import (
	"database/sql"
	"gatsby/sqlutils"
	"reflect"
)

const DefaultSliceCap = 200

/*
Scan data from sql.Rows and returns a map slice.

This returns []map[string]interface{}
*/
func CreateStructSliceFromRows(val PtrRecord, rows *sql.Rows) (interface{}, error) {
	var value = reflect.Indirect(reflect.ValueOf(val))
	var typeOfVal = value.Type()
	var sliceOfVal = reflect.SliceOf(typeOfVal)
	var slice = reflect.MakeSlice(sliceOfVal, 0, DefaultSliceCap)
	var err error
	for rows.Next() {
		var newValue = reflect.New(typeOfVal)
		if err = FillFromRows(newValue.Interface(), rows); err != nil {
			return slice.Interface(), err
		}
		slice = reflect.Append(slice, reflect.Indirect(newValue))
	}
	if err = rows.Err(); err != nil {
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
	return db.Query(sqlutils.BuildSelectClause(val))
}

func QuerySelectWith(db *sql.DB, val PtrRecord, postSql string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(sqlutils.BuildSelectClause(val)+" "+postSql, args...)
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
	var whereSql, args = sqlutils.BuildWhereClauseWithAndOp(conds)
	var sql = sqlutils.BuildSelectClause(val) + whereSql
	var rows, err = db.Query(sql, args...)
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
	var rows, err = db.Query(sql, args...)
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
