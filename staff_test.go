package gatsby

import "testing"

func TestStaffCRUD(t *testing.T) {
	var db = openDB()

	SetupConnection(db, DriverPg)

	var staff = Staff{Name: "NameA"}

	var manager = EntityManager{}
	res := resultSuccess(t, manager.Create(&staff))
	resultSuccess(t, staff.Load(res.Id))

	staff.Name = "NameB"
	resultSuccess(t, staff.Update())

	resultSuccess(t, staff.Delete())
	// resultSuccess(t, manager.Update(&staff))
}
