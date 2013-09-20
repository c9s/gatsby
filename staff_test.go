package gatsby

import "testing"

func BenchmarkCreateStaff(b *testing.B) {
	var db = openDB()
	SetupConnection(db, DriverPg)
	var staff = NewRecord(&Staff{Name: "Foo"}).(*Staff)
	for i := 0; i < b.N; i++ {
		staff.Create()
	}
}

func TestStaffCRUD(t *testing.T) {
	var db = openDB()

	SetupConnection(db, DriverPg)

	var staff = NewRecord(&Staff{Name: "Foo"}).(*Staff)
	resultSuccess(t, staff.Create())

	resultSuccess(t, staff.LoadByCols(WhereMap{
		"name": "foo",
	}))

	staff.Name = "NameB"
	resultSuccess(t, staff.Update())

	resultSuccess(t, staff.Delete())
}
