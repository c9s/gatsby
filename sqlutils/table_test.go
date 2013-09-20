package sqlutils

import "testing"

func BenchmarkGetTableName(b *testing.B) {
	s := Staff{}
	for i := 0; i < b.N; i++ {
		GetTableName(&s)
	}
}
