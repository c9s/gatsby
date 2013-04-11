package sqlutils
import "database/sql"
import "errors"

func Delete(db *sql.DB, val interface{}) (*Result) {
	sql := "DELETE FROM " + GetTableName(val) + " WHERE id = $1"

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
	return r
}

