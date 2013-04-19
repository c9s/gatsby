package sqlfragments

import "testing"

func TestFragments(t *testing.T) {
	fragment := New()
	fragment.AppendQuery("name = ?", "John")
	fragment.AppendQuery("phone = ?", "John")
	sql := fragment.Join("OR")
	if sql != "" {
		t.Fatal("sqlfragments fatal error")
	}
}


