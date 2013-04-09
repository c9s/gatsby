package sqlutils
import "testing"

type FooRecord struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Internal int `json:-`
}

func TestColumnNamesParsing(t * testing.T) {
	var columns []string
	columns = ParseColumnNames( FooRecord{Id:3, Name: "Mary"} )

	if len(columns) == 0 {
		t.Fail()
	}

	columns = ParseColumnNames( FooRecord{Id:4, Name: "John"} )
	if len(columns) == 0 {
		t.Fail()
	}
}

