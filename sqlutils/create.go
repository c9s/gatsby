package sqlutils
import "database/sql"
// import "fmt"


// id, err := sqlutils.Create(struct pointer)
func Create(db *sql.DB, val interface{}, driver int) (*Result) {
	sql , args := BuildInsertClause(val)

	err := CheckRequired(val)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	result := NewResult(sql)

	// get the autoincrement id from result
	if driver == DriverPg {
		col := GetPrimaryKeyColumnName(val)
		sql = sql + " RETURNING " + *col
		rows, err := PrepareAndQuery(db,sql,args...)

		defer func() { rows.Close() }()

		if err != nil {
			return NewErrorResult(err,sql)
		}
		id, err := GetReturningIdFromRows(rows)
		if err != nil {
			return NewErrorResult(err,sql)
		}
		if val.(PrimaryKey) != nil {
			val.(PrimaryKey).SetPkId(id)
		}
		result.Id = id
	} else if driver == DriverMysql {
		res, err := db.Exec(sql,args...)
		if err != nil {
			return NewErrorResult(err,sql)
		}
		result.Id, err = res.LastInsertId()
		if err != nil {
			return NewErrorResult(err,sql)
		}
	} else {
		panic("Unsupported driver type")
	}
	return result
}


