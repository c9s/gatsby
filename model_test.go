package gatsby
import "testing"

func TestModelSelect(t *testing.T) {
	model := NewModel("staffs")
	model.Select("id","name","columns").WhereFromMap(map[string]interface{} {
		"name": "John",
	})
	sql := model.String()

	t.Log(sql)
	if sql == "" {
		t.Fatal("SQL Select Error")
	}
}


func TestModelInsert(t *testing.T) {
	model := NewModel("staffs")
	model.Insert(map[string]interface{} {
		"name": "John",
	})
	sql := model.String()

	t.Log(sql)
	if sql == "" {
		t.Fatal("Insert SQL Error")
	}
}


func TestModelUpdate(t *testing.T) {
	model := NewModel("staffs")
	model.Update(map[string]interface{} {
		"name": "John",
	})
	model.WhereFromMap(map[string]interface{} {
		"id": 3,
	})
	sql := model.String()

	t.Log(sql)
	if sql == "" {
		t.Fatal("Update SQL Error")
	}
}

