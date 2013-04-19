package gatsby

import "testing"

func TestQuerySelect(t *testing.T) {
	query := NewQuery("staffs")
	query.Select("id", "name", "columns").WhereFromMap(map[string]interface{}{
		"name": "John",
	})
	sql := query.String()

	t.Log(sql)
	if sql == "" {
		t.Fatal("SQL Select Error")
	}
}

func TestQueryInsert(t *testing.T) {
	query := NewQuery("staffs")
	query.Insert(map[string]interface{}{
		"name": "John",
	})
	sql := query.String()

	t.Log(sql)
	if sql == "" {
		t.Fatal("Insert SQL Error")
	}
}

func TestQueryUpdate(t *testing.T) {
	query := NewQuery("staffs")
	query.Update(map[string]interface{}{
		"name": "John",
	})
	query.WhereFromMap(map[string]interface{}{
		"id": 3,
	})
	sql := query.String()

	t.Log(sql)
	if sql == "" {
		t.Fatal("Update SQL Error")
	}
}
