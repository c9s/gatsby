package sqlutils
import "database/sql"
import "errors"

func Delete(db *sql.DB, val interface{}) (*Result) {
	pkName := GetPrimaryKeyColumnName(val)

	if pkName == nil {
		return NewErrorResult( errors.New("PrimaryKey column is not defined."),"")
	}

	sql := "DELETE FROM " + GetTableName(val) + " WHERE " + *pkName + " = $1"

	if val.(PrimaryKey) == nil {
		return NewErrorResult(errors.New("PrimaryKey interface is required."),sql)
	}

	id := val.(PrimaryKey).GetPkId()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	r := NewResult(sql)
	r.Result = res
	r.Id = id
	return r
}

