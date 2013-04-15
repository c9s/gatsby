package sqlutils
import "testing"


func TestLoad(t *testing.T) {
	var db = openDB()

	staff := Staff{Name: "John", Gender: "m", Phone: "1234567"}

	r := Create(db, &staff, DriverPg)
	if r.Error != nil {
		t.Fatal(r.Error)
	}
	if r.Id == -1 {
		t.Fatal("Can not create record")
	}
	t.Logf("staff id: %d", r.Id)





	staff2 := Staff{}
	r = Load(db,&staff2, r.Id)
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
	res := LoadByCols(db, &staff3, map[string]interface{} {
		"Phone": "1234567",
	})
	if res.Error != nil {
		t.Fatal(res.Error)
	}
	if staff3.Id == 0 {
		t.Fatal(res.Error)
	}
	if staff3.Phone != "1234567" {
		t.Fatal(res.Error)
	}


	r = Delete(db,&staff)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

}
