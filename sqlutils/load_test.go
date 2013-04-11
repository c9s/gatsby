package sqlutils
import "testing"

func TestLoad(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	staff := Staff{Name: "John", Gender: "m", Phone: "1234567"}

	id, err := Create(db,&staff)
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal("Can not create record")
	}
	t.Logf("staff id: %d", id)

	staff2 := Staff{}
	err = Load(db,&staff2, id)

	if err != nil {
		t.Fatal(err)
	}

	if staff2.Id == 0 {
		t.Fatal("Can not load record")
	}


}
