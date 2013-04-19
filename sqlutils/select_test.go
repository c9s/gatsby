package sqlutils
import "testing"
import "strings"


func TestBuildSelectColumns(t * testing.T) {
	str := BuildSelectColumnClauseFromStruct( &fooRecord{Id:4, Name: "John"} )
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
	sql := BuildSelectClauseFromStruct(&staff)
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


