package sqlutils

import "testing"
import "sort"

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
	foo := fooRecord{Id: 1}
	v := GetPrimaryKeyValue(&foo)
	if v == nil {
		t.Fatal("Primary key value not found.")
	}
	if *v != 1 {
		t.Fatal("Unexpected primary key value.")
	}
	t.Logf("Primary key value: %d", *v)

	SetPrimaryKeyValue(&foo, 2)
	v = GetPrimaryKeyValue(&foo)
	if *v != 2 {
		t.Fatal("Primary key value is not updated.")
	}
}

func TestPrimaryKeyColumnValueFound2(t *testing.T) {
	v := GetPrimaryKeyValue(&fooRecord{})
	if v == nil {
		t.Fatal("Primary key value not found.")
	}
	t.Logf("Primary key value: %d", *v)
}

func TestColumnNameMap(t *testing.T) {
	columns := GetColumnValueMap(&fooRecord{Id: 3, Name: "Mary"})
	t.Log(columns)
	if len(columns) == 0 {
		t.Fail()
	}
}

func TestColumnNamesParsing(t *testing.T) {
	columns := ReflectColumnNames(&fooRecord{Id: 3, Name: "Mary"})

	// sort.Strings(columns)
	t.Log(columns)

	i := sort.SearchStrings(columns, "Internal")
	if columns[i] == "Internal" {
		t.Fail()
	}

	if len(columns) != 3 {
		t.Fail()
	}

	columns = ReflectColumnNames(&fooRecord{Id: 4, Name: "John"})
	t.Log(columns)
	if len(columns) != 3 {
		t.Fail()
	}
}

func BenchmarkGetPrimaryKeyValue(b *testing.B) {
	foo := fooRecord{Id: 3, Name: "Mary"}
	for i := 0; i < b.N; i++ {
		GetPrimaryKeyValue(&foo)
	}
}

func BenchmarkPrimaryKeyColumnName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimaryKeyColumnName(&fooRecord{})
	}
}

func BenchmarkGetColumnValueMap(b *testing.B) {
	foo := fooRecord{Id: 3, Name: "Mary"}
	for i := 0; i < b.N; i++ {
		GetColumnValueMap(&foo)
	}
}
