package gatsby

import "database/sql"
import "fmt"
import "gatsby/sqlutils"

// Load record by primary key value.
func Load(db *sql.DB, val interface{}, pkId int64) *Result {
	pName := sqlutils.GetPrimaryKeyColumnName(val)
	if pName == nil {
		panic("primary key is required.")
	}
	sqlstring := sqlutils.BuildSelectClause(val) + fmt.Sprintf(" WHERE %s = $1 LIMIT 1", *pName)
	rows, err := db.Query(sqlstring, pkId)
	if err != nil {
		return NewErrorResult(err, sqlstring)
	}

	defer func() { rows.Close() }()

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

func LoadByCols(db *sql.DB, val interface{}, cols map[string]interface{}) *Result {
	sqlstring := sqlutils.BuildSelectClause(val)
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(cols)

	sqlstring += whereSql

	rows, err := db.Query(sqlstring, args...)
	if err != nil {
		return NewErrorResult(err, sqlstring)
	}

	defer func() { rows.Close() }()

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
