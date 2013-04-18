package sqlutils
import "database/sql"


func PrepareAndQuery(db *sql.DB, sql string, args ...interface{}) (*sql.Rows,error) {
	/*
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer func() { stmt.Close() }()
	*/
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

