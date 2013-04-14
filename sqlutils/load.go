package sqlutils
import "database/sql"
import "fmt"

// Load record by primary key value.
func Load(db *sql.DB, val interface{}, pkId int64) (*Result) {
	pName := GetPrimaryKeyColumnName(val)
	if pName == nil {
		panic("primary key is required.")
	}
	sql := BuildSelectClause(val) + fmt.Sprintf(" WHERE %s = $1 LIMIT 1", *pName)
	rows, err := PrepareAndQuery(db, sql, pkId)
	rows.Next()
	err = FillFromRow(val,rows)
	if err != nil {
		return NewErrorResult(err,sql)
	}
	return NewResult(sql)
}

func LoadByCols(db *sql.DB, val interface{}, cols map[string] interface{}) (*Result) {
	sql := BuildSelectClause(val)
	whereSql, args := BuildWhereClauseWithAndOp(cols)

	sql += whereSql

	rows, err := PrepareAndQuery(db, sql, args...)
	rows.Next()
	err = FillFromRow(val,rows)

	if err != nil {
		return NewErrorResult(err,sql)
	}
	return NewResult(sql)
}
