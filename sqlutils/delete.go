package sqlutils
import "database/sql"

func Delete(db *sql.DB, val interface{}) (sql.Result, error) {
	sql := "DELETE FROM " + GetTableName(val) + " WHERE id = $1"

	if val.(PrimaryKey) == nil {
		panic("PrimaryKey interface is required.")
	}


	id := val.(PrimaryKey).GetPkId()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

