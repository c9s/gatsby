package gatsby

import (
	"gatsby/sqlutils"
)

// import "fmt"

// id, err := Create(db pointer, struct pointer)
func Create(e interface{}, val interface{}, driver int) *Result {
	var err error
	var executor, ok = e.(Executor)

	if !ok {
		panic("Not an Executor type")
	}

	err = sqlutils.CheckRequired(val)
	if err != nil {
		return NewErrorResult(err, "")
	}

	var sqlStr, args = sqlutils.BuildInsertClause(val)
	result := NewResult(sqlStr)

	// get the autoincrement id from result
	if driver == DriverPg {
		col := sqlutils.GetPrimaryKeyColumnName(val)
		sqlStr = sqlStr + " RETURNING " + *col
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
		if val.(sqlutils.PrimaryKey) != nil {
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

		if val.(sqlutils.PrimaryKey) != nil {
			val.(sqlutils.PrimaryKey).SetPkId(result.Id)
		} else {
			sqlutils.SetPrimaryKeyValue(val, result.Id)
		}
	} else {
		panic("Unsupported driver type")
	}
	return result
}
