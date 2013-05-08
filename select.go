package gatsby

import "database/sql"
import "gatsby/sqlutils"

func Select(db *sql.DB, val interface{}) (interface{}, *Result) {
	sql := sqlutils.BuildSelectClause(val)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}
	slice, err := sqlutils.CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}

func QuerySelect(db *sql.DB, val interface{}) (*sql.Rows, error) {
	sql := sqlutils.BuildSelectClause(val)
	return db.Query(sql)
}

func QuerySelectWith(db *sql.DB, val interface{}, postSql string, args ...interface{}) (*sql.Rows, error) {
	sql := sqlutils.BuildSelectClause(val) + " " + postSql
	return db.Query(sql, args...)
}

// Select a table and returns objects
func SelectWith(db *sql.DB, val interface{}, postSql string, args ...interface{}) (interface{}, *Result) {
	sql := sqlutils.BuildSelectClause(val) + " " + postSql
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}

	slice, err := sqlutils.CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}

func SelectWhere(db *sql.DB, val interface{}, conds map[string]interface{}) (interface{}, *Result) {
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(conds)
	sql := sqlutils.BuildSelectClause(val) + whereSql
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}

	slice, err := sqlutils.CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}

func SelectFromQuery(db *sql.DB, val interface{}, sql string, args ...interface{}) (interface{}, *Result) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, NewErrorResult(err, sql)
	}
	slice, err := sqlutils.CreateStructSliceFromRows(val, rows)
	if err != nil {
		return slice, NewErrorResult(err, sql)
	}
	return slice, NewResult(sql)
}
