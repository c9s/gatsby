package gatsby

import "testing"

func resultSuccess(t *testing.T, res *Result) *Result {
	if res.Error != nil {
		if res.Sql != "" {
			t.Log(res.Sql)
		}
		t.Fatal(res.Error)
	}
	return res
}

func TestLoad(t *testing.T) {
	var db = openDB()

	staff := Staff{Name: "John", Gender: "m", Phone: "1234567"}

	r := resultSuccess(t, Create(db, &staff, DriverPg))
	t.Logf("staff id: %d", r.Id)

	staff2 := Staff{}
	r = Load(db, &staff2, r.Id)
	t.Log(r.Sql)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	if staff2.Id == 0 {
		t.Fatal("Can not load record")
	}
	if staff2.Phone != "1234567" {
		t.Fatal("Can not load record")
	}

	staff3 := Staff{}
	res := resultSuccess(t, LoadByCols(db, &staff3, map[string]interface{}{
		"Phone": "1234567",
	}))
	if staff3.Id == 0 {
		t.Fatal(res.Error)
	}
	if staff3.Phone != "1234567" {
		t.Fatal(res.Error)
	}

	staff4 := Staff{}
	res = resultSuccess(t, LoadByCols(db, &staff4, WhereMap{
		"Phone": "1234567",
	}))
	if staff4.Id == 0 {
		t.Fatal(res.Error)
	}
	if staff4.Phone != "1234567" {
		t.Fatal(res.Error)
	}

	resultSuccess(t, Delete(db, &staff))
}
