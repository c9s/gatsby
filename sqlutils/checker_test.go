package sqlutils
import "testing"

func TestRequireChecking2(t *testing.T) {
	staff := Staff{Id:4, Name: "John", Gender: "m", Phone: "0975277696"}
	err := CheckRequired(&staff)
	if err != nil {
		t.Fatal("Name field is required.")
	}
}

func TestRequireChecking(t *testing.T) {
	staff := Staff{Id:4, Gender: "m", Phone: "0975277696"}
	err := CheckRequired(&staff)
	if err == nil {
		t.Fatal("Name field should be required.")
	}
}



