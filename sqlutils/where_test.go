package sqlutils
import "testing"
import "strings"

func TestBuildWhereClause(t *testing.T) {
	argMap := map[string]interface{} {
		"name": "foo",
		"id": 123,
	}
	sql, args := BuildWhereClauseWithAndOp(argMap)

	if strings.Index(sql, "name = $") == -1 {
		t.Fatal(sql)
	}
	if strings.Index(sql, "id = $") == -1 {
		t.Fatal(sql)
	}
	_ = args

	/*
	if args[0] != "foo" {
		t.Fatal("not foo")
	}
	if args[1] != 123 {
		t.Fatal("not 123")
	}
	*/
}


