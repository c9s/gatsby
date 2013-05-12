package gatsby

import "gatsby/sqlutils"
import "fmt"
import "errors"

func Update(executor Executor, val interface{}) *Result {
	pkName := sqlutils.GetPrimaryKeyColumnName(val)
	if pkName == nil {
		panic("primary key column is not defined.")
	}

	sql, values := sqlutils.BuildUpdateClause(val)

	if _, ok := val.(sqlutils.PrimaryKey); ok {
		var id = val.(sqlutils.PrimaryKey).GetPkId()
		values = append(values, id)
	} else {
		var id = sqlutils.GetPrimaryKeyValue(val)
		if id == nil {
			return NewErrorResult(errors.New("primary key field is required."), "")
		}
		values = append(values, *id)
	}

	sql += fmt.Sprintf(" WHERE %s = $%d", *pkName, len(values))

	stmt, err := executor.Prepare(sql)
	if err != nil {
		return NewErrorResult(err, sql)
	}

	defer stmt.Close()

	res, err := stmt.Exec(values...)
	if err != nil {
		return NewErrorResult(err, sql)
	}

	result := NewResult(sql)
	result.Result = res
	return result
}
