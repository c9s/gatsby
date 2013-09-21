package sqlutils

import "testing"

func TestBuildInsertClause(t *testing.T) {
	foo := fooRecord{Id: 3, Name: "Mary"}
	sql, args := BuildInsertClause(&foo)

	if len(sql) == 0 {
		t.Fatal("Empty SQL")
	}
	if len(args) == 0 {
		t.Fatal("Empty argument")
	}
}

func BenchmarkBuildInsertClause(b *testing.B) {
	foo := fooRecord{Id: 3, Name: "Mary"}
	for i := 0; i < b.N; i++ {
		BuildInsertClause(&foo)
	}
}
