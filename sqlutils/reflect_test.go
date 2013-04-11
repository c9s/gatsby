package sqlutils
import "testing"
import "sort"

type fooRecord struct {
	Id       int64  `json:"id" field:",primary"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Internal int    `json:-`
}


func TestTableName(t *testing.T) {
	n := GetTableName(&fooRecord{})
	if n != "foo_records" {
		t.Fatal("table name is not correct: " + n)
	}
}

func TestPrimaryKeyColumnName(t *testing.T) {
	n := GetPrimaryKeyColumnName(&fooRecord{})
	if n == nil {
		t.Fatal("Primary key column not found.")
	}
}

func TestPrimaryKeyColumnValueFound(t *testing.T) {
	v := GetPrimaryKeyValue(&fooRecord{ Id: 1 })
	if v == nil {
		t.Fatal("Primary key value not found.")
	}
	if *v != 1 {
		t.Fatal("Unexpected primary key value.")
	}
	t.Logf("Primary key value: %d", *v)
}

func TestPrimaryKeyColumnValueFound2(t *testing.T) {
	v := GetPrimaryKeyValue(&fooRecord{})
	if v == nil {
		t.Fatal("Primary key value not found.")
	}
	t.Logf("Primary key value: %d", *v)
}

func TestColumnNameMap(t *testing.T) {
	columns := GetColumnValueMap( &fooRecord{ Id: 3, Name: "Mary" } )
	t.Log(columns)
	if len(columns) == 0 {
		t.Fail()
	}
}


func TestColumnNamesParsing(t * testing.T) {
	columns := ParseColumnNames( &fooRecord{Id:3, Name: "Mary"} )

	// sort.Strings(columns)
	t.Log(columns)

	i := sort.SearchStrings(columns, "Internal")
	if columns[i] == "Internal" {
		t.Fail()
	}

	if len(columns) != 3 {
		t.Fail()
	}

	columns = ParseColumnNames( &fooRecord{Id:4, Name: "John"} )
	t.Log(columns)
	if len(columns) != 3 {
		t.Fail()
	}
}

