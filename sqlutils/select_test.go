package sqlutils
import "testing"
import "strings"

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

