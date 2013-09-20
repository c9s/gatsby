package sqlutils

import (
	"strings"
	"testing"
)

func BenchmarkBuildSelectColumnClauseFromStructWithAlias(b *testing.B) {
	foo := fooRecord{}
	for i := 0; i < b.N; i++ {
		BuildSelectColumnClauseFromStructWithAlias(&foo, "foo")
	}
}

func BenchmarkBuildSelectColumnClauseFromStruct(b *testing.B) {
	foo := fooRecord{Id: 4, Name: "John"}
	for i := 0; i < b.N; i++ {
		BuildSelectColumnClauseFromStruct(&foo)
	}
}

func TestBuildSelectClauseWithAlias(t *testing.T) {
	str := BuildSelectColumnClauseFromStructWithAlias(&fooRecord{}, "foo")
	t.Log(str)
}

func TestBuildSelectColumns(t *testing.T) {
	str := BuildSelectColumnClauseFromStruct(&fooRecord{Id: 4, Name: "John"})
	if len(str) == 0 {
		t.Fail()
	}
	if !strings.Contains(str, "id, name, type") {
		t.Fatal(str)
	}
	t.Log(str)
}

func TestBuildSelectClause(t *testing.T) {
	staff := Staff{Id: 4, Name: "John", Gender: "m", Phone: "0975277696"}
	sql := BuildSelectClause(&staff)
	if !strings.Contains(sql, "SELECT id, name, gender, staff_type, phone, birthday") {
		t.Fatal(sql)
	}
	if !strings.Contains(sql, "FROM staffs") {
		t.Fatal(sql)
	}
}

func chkResult(t *testing.T, res *Result) {
	if res.Error != nil {
		t.Fatal(res.Error, res.Sql)
	}
}
