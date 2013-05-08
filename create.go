package gatsby

import "gatsby/sqlutils"

// import "fmt"

// id, err := sqlutils.Create(struct pointer)
func Create(e interface{}, val interface{}, driver int) *Result {
	var executor, ok = e.(Executor)

	if !ok {
		panic("Not an Executor type")
	}

	var sqlStr, args = sqlutils.BuildInsertClause(val)

	err := sqlutils.CheckRequired(val)
	if err != nil {
		return NewErrorResult(err, sqlStr)
	}

	result := NewResult(sqlStr)

	// get the autoincrement id from result
	if driver == DriverPg {
		col := sqlutils.GetPrimaryKeyColumnName(val)
		sqlStr = sqlStr + " RETURNING " + *col
		rows, err := executor.Query(sqlStr, args...)

		defer func() { rows.Close() }()

		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		id, err := sqlutils.GetPgReturningIdFromRows(rows)
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		if val.(sqlutils.PrimaryKey) != nil {
			val.(sqlutils.PrimaryKey).SetPkId(id)
		}
		result.Id = id
	} else if driver == DriverMysql {
		res, err := executor.Exec(sqlStr, args...)
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		result.Id, err = res.LastInsertId()
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
	} else {
		panic("Unsupported driver type")
	}
	return result
}
