package gatsby
import "database/sql"
import "errors"
import "gatsby/sqlutils"

func Delete(db *sql.DB, val interface{}) (*Result) {
	pkName := sqlutils.GetPrimaryKeyColumnName(val)

	if pkName == nil {
		return NewErrorResult( errors.New("PrimaryKey column is not defined."),"")
	}

	sql := "DELETE FROM " + sqlutils.GetTableName(val) + " WHERE " + *pkName + " = $1"

	if val.(sqlutils.PrimaryKey) == nil {
		return NewErrorResult(errors.New("PrimaryKey interface is required."),sql)
	}

	id := val.(sqlutils.PrimaryKey).GetPkId()
	res, err := db.Exec(sql, id)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	r := NewResult(sql)
	r.Result = res
	r.Id = id
	return r
}

