package sqlutils
import "testing"

func TestBuildWhereClause(t *testing.T) {
	argMap := map[string]interface{} {
		"name": "foo",
		"id": 123,
	}
	sql, args := BuildWhereClauseWithAndOp(argMap)
	if sql != " WHERE name = $1 AND id = $2" {
		t.Fatal(sql)
	}
	if args[0] != "foo" {
		t.Fail()
	}
	if args[1] != 123 {
		t.Fail()
	}
}


