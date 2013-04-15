package sqlutils
import "testing"

func TestBuildSelectColumns(t * testing.T) {
	str := BuildSelectColumnClause( &fooRecord{Id:4, Name: "John"} )
	if len(str) == 0 {
		t.Fail()
	}
	if str != "id,name,type" {
		t.Fatal(str)
	}
}

func TestBuildSelectClause(t * testing.T) {
	staff := Staff{Id:4, Name: "John", Gender: "m", Phone: "0975277696"}
	sql := BuildSelectClause(&staff)
	if sql != "SELECT id,name,gender,staff_type,phone,birthday FROM staffs" {
		t.Fatal(sql)
	}
}

