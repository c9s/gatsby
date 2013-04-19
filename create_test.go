package gatsby

import "testing"

func TestCreate(t *testing.T) {
	var db = openDB()
	staff := Staff{}
	staff.Name = "John"
	Create(db, &staff, DriverPg)
	Delete(db, &staff)
}
