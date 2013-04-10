package sqlutils
import "testing"

func TestBuildInsertClause(t *testing.T) {
	foo := fooRecord{ Id: 3, Name: "Mary" }
	t.Log( BuildInsertColumnClause(foo) )
}

