package sqlutils
import "testing"

func TestBuildInsertClause(t *testing.T) {
	foo := fooRecord{ Id: 3, Name: "Mary" }
	sql, args := BuildInsertColumnClause(&foo)
	_ = sql
	_ = args
}


