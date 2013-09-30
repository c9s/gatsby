package gatsby

import (
	"database/sql"
	"github.com/c9s/gatsby/sqlutils"
)

type WhereMap map[string]interface{}

// Load record by primary key value.
var loadQueryCache = map[string]string{}

/*
Fill a record object by executing QueryRow from sql.DB object,
this method is faster than the DB.Query method.
*/
func LoadFromQueryRow(db *sql.DB, val PtrRecord, sqlstring string, args ...interface{}) *Result {
	var err error
	var row = db.QueryRow(sqlstring, args...)
	if err = FillFromRows(val, row); err != nil {
		if err == sql.ErrNoRows {
			res := NewResult(sqlstring)
			res.IsEmpty = true
			return res
		}
		return NewErrorResult(err, sqlstring)
	}
	return NewResult(sqlstring)
}

/*
Fill a record object by executing a SQL query.
*/
func LoadFromQuery(db *sql.DB, val PtrRecord, sqlstring string, args ...interface{}) *Result {
	var err error
	row := db.QueryRow(sqlstring, args...)
	if err = FillFromRows(val, row); err != nil {
		if err == sql.ErrNoRows {
			res := NewResult(sqlstring)
			res.IsEmpty = true
			return res
		}
		return NewErrorResult(err, sqlstring)
	}
	return NewResult(sqlstring)
}

func LoadWith(db *sql.DB, val PtrRecord, postQuery string, args ...interface{}) *Result {
	var sqlstring = sqlutils.BuildSelectClause(val) + " " + postQuery + sqlutils.BuildLimitClause(1)
	return LoadFromQueryRow(db, val, sqlstring, args...)
}

func Load(db *sql.DB, val PtrRecord, pkId int64) *Result {
	var sqlstring = sqlutils.BuildLoadClause(val)
	return LoadFromQueryRow(db, val, sqlstring, pkId)
}

/*
Load record from a where condition map
*/
func LoadByCols(db *sql.DB, val PtrRecord, cols WhereMap) *Result {
	var sqlstring = sqlutils.BuildSelectClause(val)
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(cols)
	sqlstring += whereSql + sqlutils.BuildLimitClause(1)
	return LoadFromQueryRow(db, val, sqlstring, args...)
}
