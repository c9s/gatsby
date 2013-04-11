package sqlutils
import "testing"
import "sort"
import "database/sql"
import _ "github.com/bmizerany/pq"

type fooRecord struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Internal int    `json:-`
}

func TestColumnNameMap(t *testing.T) {
	columns := GetColumnMap( fooRecord{ Id: 3, Name: "Mary" } )

	t.Log(columns)

	if len(columns) == 0 {
		t.Fail()
	}
}


func TestColumnNamesParsing(t * testing.T) {
	columns := ParseColumnNames( fooRecord{Id:3, Name: "Mary"} )

	// sort.Strings(columns)
	t.Log(columns)

	i := sort.SearchStrings(columns, "Internal")
	if columns[i] == "Internal" {
		t.Fail()
	}

	if len(columns) != 3 {
		t.Fail()
	}

	columns = ParseColumnNames( fooRecord{Id:4, Name: "John"} )
	t.Log(columns)
	if len(columns) != 3 {
		t.Fail()
	}
}


func TestBuildSelectColumns(t * testing.T) {
	str := BuildSelectColumnClause( fooRecord{Id:4, Name: "John"} )
	if len(str) == 0 {
		t.Fail()
	}
	if str != "id,name,type" {
		t.Fatal(str)
	}
}



type Staff struct {
	Id        int `json:"id"`
	Name      string `json:"name" field:",required"`
	Gender    string `json:"gender"`
	StaffType string `json:"staff_type"` // valid types: doctor, nurse, ...etc
	Phone     string `json:"phone"`
}

func TestBuildSelectClause(t * testing.T) {
	staff := Staff{Id:4, Name: "John", Gender: "m", Phone: "0975277696"}
	sql := BuildSelectClause(staff)
	if sql != "SELECT id,name,gender,staff_type,phone FROM staffs" {
		t.Fatal(sql)
	}
}

func TestRequireChecking2(t *testing.T) {
	staff := Staff{Id:4, Name: "John", Gender: "m", Phone: "0975277696"}
	err := CheckRequired(&staff)
	if err != nil {
		t.Fatal("Name field is required.")
	}
}

func TestRequireChecking(t *testing.T) {
	staff := Staff{Id:4, Gender: "m", Phone: "0975277696"}
	err := CheckRequired(&staff)
	if err == nil {
		t.Fatal("Name field should be required.")
	}
}



func TestFillRecord(t * testing.T) {
	staff := Staff{}
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=drshine_itsystem sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	sql := BuildSelectClause(staff) + " WHERE id = $1"
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


