package gatsby

import "database/sql"
import "gatsby/sqlutils"

// import "fmt"

// id, err := sqlutils.Create(struct pointer)
func Create(db *sql.DB, val interface{}, driver int) *Result {
	sql, args := sqlutils.BuildInsertClause(val)

	err := sqlutils.CheckRequired(val)
	if err != nil {
		return NewErrorResult(err, sql)
	}

	result := NewResult(sql)

	// get the autoincrement id from result
	if driver == DriverPg {
		col := sqlutils.GetPrimaryKeyColumnName(val)
		sql = sql + " RETURNING " + *col
		rows, err := db.Query(sql, args...)

		defer func() { rows.Close() }()

		if err != nil {
			return NewErrorResult(err, sql)
		}
		id, err := sqlutils.GetPgReturningIdFromRows(rows)
		if err != nil {
			return NewErrorResult(err, sql)
		}
		if val.(sqlutils.PrimaryKey) != nil {
			val.(sqlutils.PrimaryKey).SetPkId(id)
		}
		result.Id = id
	} else if driver == DriverMysql {
		res, err := db.Exec(sql, args...)
		if err != nil {
			return NewErrorResult(err, sql)
		}
		result.Id, err = res.LastInsertId()
		if err != nil {
			return NewErrorResult(err, sql)
		}
	} else {
		panic("Unsupported driver type")
	}
	return result
}
