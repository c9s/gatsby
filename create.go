package gatsby

import (
	"gatsby/sqlutils"
)

// import "fmt"

// id, err := Create(db pointer, struct pointer)
func Create(executor Executor, val interface{}, driver int) *Result {
	var err error

	if err = sqlutils.CheckRequired(val); err != nil {
		return NewErrorResult(err, "")
	}

	var sqlStr, args = sqlutils.BuildInsertClause(val)
	result := NewResult(sqlStr)

	// get the autoincrement id from result
	if driver == DriverPg {
		if col := sqlutils.GetPrimaryKeyColumnName(val); col != nil {
			sqlStr = sqlStr + " RETURNING " + *col
		}
		rows, err := executor.Query(sqlStr, args...)

		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		defer rows.Close()

		id, err := sqlutils.GetPgReturningIdFromRows(rows)
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}

		// if the struct supports the primary key interface, we can set the value faster.
		result.Id = id
		if _, ok := val.(sqlutils.PrimaryKey); ok {
			val.(sqlutils.PrimaryKey).SetPkId(id)
		} else {
			sqlutils.SetPrimaryKeyValue(val, id)
		}
	} else if driver == DriverMysql {
		res, err := executor.Exec(sqlStr, args...)
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		result.Id, err = res.LastInsertId()
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}

		if _, ok := val.(sqlutils.PrimaryKey); ok {
			val.(sqlutils.PrimaryKey).SetPkId(result.Id)
		} else {
			sqlutils.SetPrimaryKeyValue(val, result.Id)
		}
	} else {
		panic("Unsupported driver type")
	}
	return result
}
