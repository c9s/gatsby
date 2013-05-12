package gatsby

import "testing"

func TestStaffCRUD(t *testing.T) {
	var db = openDB()

	SetupConnection(db, DriverPg)

	var staff *Staff = NewModel(&Staff{Name: "Foo"}).(*Staff)
	resultSuccess(t, staff.Create())

	staff.Name = "NameB"
	resultSuccess(t, staff.Update())

	resultSuccess(t, staff.Delete())
}
