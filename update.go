package gatsby

import "github.com/c9s/gatsby/sqlutils"
import "strconv"

func Update(executor Executor, val interface{}, driver int) *Result {
	pkName := sqlutils.GetPrimaryKeyColumnName(val)
	if pkName == nil {
		panic("primary key column is not defined.")
	}

	sql, values := sqlutils.BuildUpdateClause(val)

	var id = sqlutils.GetPrimaryKeyValue(val)
	values = append(values, id)

	// sql += fmt.Sprintf(" WHERE %s = $%d", *pkName, len(values))
	if driver == DriverPg {
		sql += " WHERE " + *pkName + " = $" + strconv.Itoa(len(values))
	} else {
		sql += " WHERE " + *pkName + " = ?"
	}

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
