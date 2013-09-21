package sqlutils

import "testing"
import "reflect"

func BenchmarkIndexOfChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexOfChar("primary,required,date", ",")
	}
}

func BenchmarkGetColumnNameFromTag(b *testing.B) {
	val := fooRecord{}
	t := reflect.ValueOf(&val).Elem()
	typeOfT := t.Type()

	var tag reflect.StructTag = typeOfT.Field(0).Tag
	for i := 0; i < b.N; i++ {
		GetColumnNameFromTag(&tag)
	}
}
