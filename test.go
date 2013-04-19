package gatsby

import "database/sql"

var db *sql.DB

func openDB() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
