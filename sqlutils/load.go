package sqlutils
import "database/sql"

// Load one record
func Load(db *sql.DB, val interface{}, pkId int64) (*Result) {
	sql := BuildSelectClause(val) + " WHERE id = $1 LIMIT 1"
	rows, err := PrepareAndQuery(db, sql, pkId)
	rows.Next()
	err = FillFromRow(val,rows)
	if err != nil {
		return NewErrorResult(err,sql)
	}
	return NewResult(sql)
}

