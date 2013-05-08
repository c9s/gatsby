package gatsby

import "gatsby/sqlutils"
import "fmt"

func Update(e Executor, val interface{}) *Result {
	var executor, ok = e.(Executor)
	if !ok {
		panic("Not an Executor type")
	}

	pkName := sqlutils.GetPrimaryKeyColumnName(val)
	if pkName == nil {
		panic("primary key column is not defined.")
	}

	sql, values := sqlutils.BuildUpdateClause(val)

	if val.(sqlutils.PrimaryKey) != nil {
		id := val.(sqlutils.PrimaryKey).GetPkId()
		values = append(values, id)
	}

	sql += fmt.Sprintf(" WHERE %s = $%d", *pkName, len(values))

	stmt, err := executor.Prepare(sql)

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
