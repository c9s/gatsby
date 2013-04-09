package sqlutils
import "testing"
import "sort"

type FooRecord struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Internal int `json:-`
}

func TestColumnNameMap(t *testing.T) {

}


func TestColumnNamesParsing(t * testing.T) {
	var columns []string
	columns = ParseColumnNames( FooRecord{Id:3, Name: "Mary"} )

	// sort.Strings(columns)
	t.Log(columns)
	i := sort.SearchStrings(columns, "Internal")
	if columns[i] == "Internal" {
		t.Fail()
	}

	if len(columns) != 3 {
		t.Fail()
	}

	columns = ParseColumnNames( FooRecord{Id:4, Name: "John"} )
	t.Log(columns)
	if len(columns) != 3 {
		t.Fail()
	}
}

