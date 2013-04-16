package sqlutils
import "testing"
import "strings"

func TestSelectQuery(t *testing.T) {
	var db = openDB()

	staff := Staff{Name: "John", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff, DriverPg) )

	rows, err := SelectQuery(db, &staff)
	if err != nil {
		t.Fatal(err)
	}

	var count = 0

	for rows.Next() {
		count++
	}
	if count == 0 {
		t.Fatal("select 0 record")
	}

	Delete(db, &staff)
}

func TestBuildSelectColumns(t * testing.T) {
	str := BuildSelectColumnClause( &fooRecord{Id:4, Name: "John"} )
	if len(str) == 0 {
		t.Fail()
	}
	if ! strings.Contains(str,"id,name,type") {
		t.Fatal(str)
	}
	t.Log(str)
}

func TestBuildSelectClause(t * testing.T) {
	staff := Staff{Id:4, Name: "John", Gender: "m", Phone: "0975277696"}
	sql := BuildSelectClause(&staff)
	if ! strings.Contains(sql,"SELECT id,name,gender,staff_type,phone,birthday") {
		t.Fatal(sql)
	}
	if ! strings.Contains(sql,"FROM staffs") {
		t.Fatal(sql)
	}
}

func chkResult(t *testing.T, res *Result) {
	if res.Error != nil {
		t.Fatal(res.Error, res.Sql)
	}
}


func TestSelectWhere(t *testing.T) {

}

func TestSelectWith(t *testing.T) {
	var db = openDB()
	staff := Staff{Name: "John", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff, DriverPg) )

	staff2 := Staff{Name: "Mary", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff2, DriverPg) )

	staff3 := Staff{Name: "Jack", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff3, DriverPg) )

	items, result := Select(db, &staff)
	chkResult(t, result)

	staffs := items.([]Staff)

	if len(staffs) == 0 {
		t.Fatal("found 0 record")
	}

	for _, s := range staffs {
		t.Log(s.Id)
		if s.Name == "" {
			t.Fatal("Empty name")
		}
		if s.Id > 0 {
			var res = Delete(db, &s)
			if res.Error != nil {
				t.Fatal(res.Error)
			}
		}
	}
	_ = staffs
}


