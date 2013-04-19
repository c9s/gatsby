package gatsby

import "database/sql"
import "fmt"
import "errors"
import "gatsby/sqlutils"

// Load record by primary key value.
func Load(db *sql.DB, val interface{}, pkId int64) *Result {
	pName := sqlutils.GetPrimaryKeyColumnName(val)
	if pName == nil {
		panic("primary key is required.")
	}
	sql := sqlutils.BuildSelectClause(val) + fmt.Sprintf(" WHERE %s = $1 LIMIT 1", *pName)
	rows, err := db.Query(sql, pkId)
	if err != nil {
		return NewErrorResult(err, sql)
	}

	defer func() { rows.Close() }()

	if rows.Next() {
		err = sqlutils.FillFromRow(val, rows)
		if err != nil {
			return NewErrorResult(err, sql)
		}
		return NewResult(sql)
	}
	err = rows.Err()
	if err != nil {
		return NewErrorResult(err, sql)
	}

	return NewErrorResult(errors.New("No result"), sql)
}

func LoadByCols(db *sql.DB, val interface{}, cols map[string]interface{}) *Result {
	sql := sqlutils.BuildSelectClause(val)
	whereSql, args := sqlutils.BuildWhereClauseWithAndOp(cols)

	sql += whereSql

	rows, err := db.Query(sql, args...)
	if err != nil {
		return NewErrorResult(err, sql)
	}

	defer func() { rows.Close() }()

	if rows.Next() {
		err = sqlutils.FillFromRow(val, rows)
		if err != nil {
			return NewErrorResult(err, sql)
		}
		return NewResult(sql)
	}
	err = rows.Err()
	if err != nil {
		return NewErrorResult(err, sql)
	}
	return NewErrorResult(errors.New("No result"), sql)
}
