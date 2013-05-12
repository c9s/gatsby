package gatsby

import "database/sql"
import "fmt"
import "gatsby/sqlutils"

func LoadFromQuery(db *sql.DB, val interface{}, sqlstring string, args ...interface{}) *Result {
	rows, err := db.Query(sqlstring, args...)
	if err != nil {
		return NewErrorResult(err, sqlstring)
	}
	defer rows.Close()

	if rows.Next() {
		err = sqlutils.FillFromRow(val, rows)
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

// Load record by primary key value.
func Load(db *sql.DB, val interface{}, pkId int64) *Result {
	pName := sqlutils.GetPrimaryKeyColumnName(val)
	if pName == nil {
		panic("primary key is required.")
	}
	sqlstring := sqlutils.BuildSelectClause(val) + fmt.Sprintf(" WHERE %s = $1", *pName)
	sqlstring += sqlutils.BuildLimitClause(1)
	return LoadFromQuery(db, val, sqlstring, pkId)
}

func LoadByCols(db *sql.DB, val interface{}, cols map[string]interface{}) *Result {
	sqlstring := sqlutils.BuildSelectClause(val)
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(cols)
	sqlstring += whereSql

	// we should only query one record
	sqlstring += sqlutils.BuildLimitClause(1)
	return LoadFromQuery(db, val, sqlstring, args...)
}
