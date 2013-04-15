package sqlutils
import "database/sql"
import "errors"

// This function fetch the returning ID from result rows
func GetReturningIdFromRows(rows * sql.Rows) (int64, error) {
	var id int64
	var err error
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return -1, err
		}
		return id, err
	}
	return -1, errors.New("No returning ID")
}

