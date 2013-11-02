package gatsby

import (
	"github.com/c9s/gatsby/sqlutils"
)

func GetHolderTypeByDriver(driver int) int {
	if driver == DriverSqlite || driver == DriverMysql {
		return sqlutils.QMARK_HOLDER
	} else {
		return sqlutils.NUMBER_HOLDER
	}
}

// id, err := Create(db pointer, struct pointer)
func Create(executor Executor, val interface{}, driver int) *Result {
	var err error

	if err = sqlutils.CheckRequired(val); err != nil {
		return NewErrorResult(err, "")
	}

	var sqlStr, args = sqlutils.BuildInsertClause(val, GetHolderTypeByDriver(driver))

	result := NewResult(sqlStr)

	// get the autoincrement id from result
	if driver == DriverPg {
		if col := sqlutils.GetPrimaryKeyColumnName(val); col != nil {
			sqlStr = sqlStr + " RETURNING " + *col
		}
		row := executor.QueryRow(sqlStr, args...)
		id, err := GetPgReturningIdFromRows(row)
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}

		// if the struct supports the primary key interface, we can set the value faster.
		result.Id = id
		sqlutils.SetPrimaryKeyValue(val, result.Id)
	} else if driver == DriverMysql || driver == DriverSqlite {
		res, err := executor.Exec(sqlStr, args...)
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		result.Id, err = res.LastInsertId()
		if err != nil {
			return NewErrorResult(err, sqlStr)
		}
		sqlutils.SetPrimaryKeyValue(val, result.Id)
	} else {
		panic("Unsupported driver type")
	}
	return result
}
