package gatsby

// This function fetch the returning ID from result rows
func GetPgReturningIdFromRows(row RowScanner) (int64, error) {
	var id int64
	var err error
	if err = row.Scan(&id); err != nil {
		return -1, err
	}
	return id, err
}
