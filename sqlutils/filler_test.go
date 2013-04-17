package sqlutils
import "testing"
import "database/sql"
import _ "github.com/c9s/pq"
import "time"
import "strings"


var db *sql.DB

func openDB() (*sql.DB) {
	if db != nil {
		return db
	}

    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}


func TestFill(t *testing.T) {
    var db = openDB()
	stmt, _ := db.Prepare("select created_on from staffs")
	rows, _ := stmt.Query()
	for rows.Next() {
		time1 := new(time.Time)
		rows.Scan(time1)
		t.Logf("Created on: %d",time1.Unix())
	}
}

func TestCreateMapFromRows(t *testing.T) {
	staff := Staff{}
    var db = openDB()

	// Create Staff
	staff.Id = 1
	staff.Name = "Mary"
	staff.Phone = "1234567"
	t1 := time.Now()
	staff.CreatedOn = &t1

	r := Create(db,&staff, DriverPg)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	if r.Id == -1 {
		t.Fatal("Primary key failed")
	}
	staff.Id = r.Id

	rows, _ := db.Query("select id, name from staffs")

	rows.Next()
	result, err := CreateMapFromRows(rows, new(int64), new(string) )
	if err != nil {
		t.Fatal( err )
	}

	if _, ok := result["id"] ; ! ok {
		t.Fatal("Can not read id")
	}

	if _, ok := result["name"] ; ! ok {
		t.Fatal("Can not read name")
	}
	t.Log( "Map", result )



	results, err := CreateMapsFromRows(rows, new(int64), new(string) )
	if err != nil {
		t.Fatal( err )
	}

	for _, r := range results {
		if _, ok := r["id"] ; ! ok {
			t.Fatal("Can not read id")
		}

		if _, ok := r["name"] ; ! ok {
			t.Fatal("Can not read name")
		}
	}

	Delete(db, &staff)
}

func TestFillRecord(t * testing.T) {
	staff := Staff{}
    var db = openDB()

	// Create Staff
	staff.Id = 1
	staff.Name = "Mary"
	staff.Phone = "1234567"
	t1 := time.Now()
	staff.CreatedOn = &t1

	r := Create(db,&staff, DriverPg)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	if r.Id == -1 {
		t.Fatal("Primary key failed")
	}
	staff.Id = r.Id


	sql := BuildSelectClause(&staff) + " WHERE id = $1"

	if ! strings.Contains(sql, "id,name,gender,staff_type,phone,birthday") {
		t.Fatal("Unexpected SQL: " + sql)
	}

	if ! strings.Contains(sql, "FROM staffs WHERE id = $1") {
		t.Fatal("Unexpected SQL: " + sql)
	}


	stmt , err := db.Prepare(sql)
	rows, err := stmt.Query( r.Id)

	if rows.Next() {
		err = FillFromRow(&staff,rows)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatal("No record found.")
	}

	r = Delete(db,&staff)
	t.Log(r)
	if r.Error != nil {
		t.Fatal(r.Error)
	}
}


