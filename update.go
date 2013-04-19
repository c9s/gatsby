package gatsby

import "database/sql"
import "gatsby/sqlutils"
import "fmt"

func Update(db *sql.DB, val interface{}) *Result {

	pkName := sqlutils.GetPrimaryKeyColumnName(&val)
	if pkName == nil {
		panic("primary key column is not defined.")
	}

	sql, values := sqlutils.BuildUpdateClause(val)

	if val.(sqlutils.PrimaryKey) != nil {
		id := val.(sqlutils.PrimaryKey).GetPkId()
		values = append(values, id)
	}

	sql += fmt.Sprintf(" WHERE %s = $%d", *pkName, len(values))

	stmt, err := db.Prepare(sql)

	defer func() { stmt.Close() }()

	if err != nil {
		return NewErrorResult(err, sql)
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		return NewErrorResult(err, sql)
	}

	result := NewResult(sql)
	result.Result = res
	return result
}
