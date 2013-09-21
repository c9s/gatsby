package sqlutils

import "database/sql"
import "errors"

// This function fetch the returning ID from result rows
func GetPgReturningIdFromRows(rows *sql.Rows) (int64, error) {
	var id int64
	var err error
	if rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return -1, err
		}
		return id, err
	}
	if err := rows.Err(); err != nil {
		return -1, err
	}
	return -1, errors.New("No returning ID")
}
