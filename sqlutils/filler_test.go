package sqlutils
import "testing"
import "database/sql"
import _ "github.com/bmizerany/pq"

func TestFillRecord(t * testing.T) {
	staff := Staff{}
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=drshine_itsystem sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	sql := BuildSelectClause(&staff) + " WHERE id = $1"
	if sql != "SELECT id,name,gender,staff_type,phone FROM staffs WHERE id = $1" {
		t.Fatal(sql)
	}

	_ = db

	stmt , err := db.Prepare(sql)
	rows, err := stmt.Query(1)
	rows.Next()
	err = FillFromRow(&staff,rows)
	if err != nil {
		t.Fatal(err)
	}
}


