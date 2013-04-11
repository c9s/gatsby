package sqlutils
import "database/sql"

// id, err := sqlutils.Create(struct pointer)
func Create(db *sql.DB, val interface{}) (*Result) {
	sql , args := BuildInsertClause(val)

	// for pgsql only
	sql += " RETURNING id"

	err := CheckRequired(val)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	rows, err := PrepareAndQuery(db,sql,args...)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	id, err := GetReturningIdFromRows(rows)

	if err != nil {
		return NewErrorResult(err,sql)
	}

	r := NewResult(sql)
	r.Id = id
	return r
}


