package sqlutils
import "database/sql"


func PrepareAndQuery(db *sql.DB, sql string, args ...interface{}) (*sql.Rows,error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}



