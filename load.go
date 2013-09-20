package gatsby

import (
	"database/sql"
	"gatsby/sqlutils"
)

type WhereMap map[string]interface{}

/*
Fill a record object by executing QueryRow from sql.DB object,
this method is faster than the DB.Query method.
*/
func LoadFromQueryRow(db *sql.DB, val PtrRecord, sqlstring string, args ...interface{}) *Result {
	var err error

	var row = db.QueryRow(sqlstring, args...)

	err = FillFromRows(val, row)
	if err != nil {
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
	rows, err := db.Query(sqlstring, args...)
	if err != nil {
		return NewErrorResult(err, sqlstring)
	}
	defer rows.Close()

	if rows.Next() {
		err = FillFromRows(val, rows)
		if err != nil {
			return NewErrorResult(err, sqlstring)
		}
		return NewResult(sqlstring)
	}
	err = rows.Err()
	if err != nil {
		return NewErrorResult(err, sqlstring)
	}

	res := NewResult(sqlstring)
	res.IsEmpty = true
	return res
}

func LoadWith(db *sql.DB, val PtrRecord, postQuery string, args ...interface{}) *Result {
	var sqlstring = sqlutils.BuildSelectClause(val) + " " + postQuery
	sqlstring += sqlutils.BuildLimitClause(1)
	return LoadFromQueryRow(db, val, sqlstring, args...)
}

// Load record by primary key value.
func Load(db *sql.DB, val PtrRecord, pkId int64) *Result {
	var pName = sqlutils.GetPrimaryKeyColumnName(val)
	if pName == nil {
		panic("primary key is required.")
	}
	var sqlstring = sqlutils.BuildSelectClause(val) + " WHERE " + *pName + " = $1" +
		sqlutils.BuildLimitClause(1)
	return LoadFromQueryRow(db, val, sqlstring, pkId)
}

/*
Load record from a where condition map
*/
func LoadByCols(db *sql.DB, val PtrRecord, cols WhereMap) *Result {
	var sqlstring = sqlutils.BuildSelectClause(val)
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(cols)
	sqlstring += whereSql

	// we should only query one record
	sqlstring += sqlutils.BuildLimitClause(1)
	return LoadFromQueryRow(db, val, sqlstring, args...)
}
