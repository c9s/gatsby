package gatsby
import "testing"

func TestModel(t *testing.T) {
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


