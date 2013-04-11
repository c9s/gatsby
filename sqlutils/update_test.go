package sqlutils
import "testing"

func TestBuildUpdateClause(t *testing.T) {
	foo := fooRecord{ Id: 3, Name: "Mary" }
	sql, values := BuildUpdateClause(&foo)
	if sql != "UPDATE foo_records SET id = $1, name = $2, type = $3" {
		t.Fatal("SQL Error: " + sql)
	}
	if len(values) != 3 {
		t.Fatal("Value length error")
	}
}

