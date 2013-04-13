package sqlutils
import "testing"

func TestBuildInsertClause(t *testing.T) {
	foo := fooRecord{ Id: 3, Name: "Mary" }
	sql, args := BuildInsertClause(&foo)

	if len(sql) == 0 {
		t.Fatal("Empty SQL")
	}
	if len(args) == 0 {
		t.Fatal("Empty argument")
	}
}


