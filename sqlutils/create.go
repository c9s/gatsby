package sqlutils
import "database/sql"
import "fmt"

// id, err := sqlutils.Create(struct pointer)
func Create(db *sql.DB, val interface{}) (int,error) {
	sql , args := BuildInsertColumnClause(val)

	// for pgsql only
	sql += " RETURNING id"

	err := CheckRequired(val)
	if err != nil {
		return -1, err
	}

	rows, err := PrepareAndQuery(db,sql,args...)
	if err != nil {
		return -1, err
	}
	return GetReturningIdFromRows(rows)
}


