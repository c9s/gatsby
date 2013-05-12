package gatsby

import "testing"

func TestStaffCRUD(t *testing.T) {
	var db = openDB()

	SetupConnection(db, DriverPg)

	baseRecord := BaseRecord{}
	baseRecord.SetTarget(&Staff{})

	/*
		var staff = Staff{Name: "NameA"}
		staff.SetTarget(&staff)
	*/

	/*
		var res = resultSuccess(t, staff.Create())
		resultSuccess(t, staff.Load(res.Id))

		staff.Name = "NameB"
		resultSuccess(t, staff.Update())
		resultSuccess(t, staff.Delete())
	*/
}
