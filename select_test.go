package gatsby

import "testing"
import "gatsby/sqlutils"

func chkResult(t *testing.T, res *Result) {
	if res.Error != nil {
		t.Fatal(res.Error, res.Sql)
	}
}

func TestSelectQueryWithAlias(t *testing.T) {
	var db = openDB()
	db.Query("delete from staffs;")

	staff := Staff{Name: "John", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff, DriverPg))
	defer Delete(db, &staff)

	var selectClause = sqlutils.BuildSelectClauseWithAlias(&staff, "s")
	t.Log(selectClause)

	rows, err := db.Query(selectClause)
	if err != nil {
		t.Fatal(err)
	}

	data, err := CreateStructSliceFromRows(&staff, rows)
	if err != nil {
		t.Fatal(err)
	}

	var staffs = data.([]Staff)
	for _, s := range staffs {
		t.Log(s)
	}
	db.Query("delete from staffs;")
}

func TestSelectQuery(t *testing.T) {
	var db = openDB()
	staff := Staff{Name: "John", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff, DriverPg))

	rows, err := QuerySelect(db, &staff)
	if err != nil {
		t.Fatal(err)
	}
	var count = 0
	for rows.Next() {
		count++
	}
	if count == 0 {
		t.Fatal("select 0 record")
	}
	Delete(db, &staff)
}

func TestSelect(t *testing.T) {
	var db = openDB()
	staff := Staff{Name: "John", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff, DriverPg))

	staff2 := Staff{Name: "Mary", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff2, DriverPg))

	staff3 := Staff{Name: "Jack", Gender: "m", Phone: "0975277696"}
	chkResult(t, Create(db, &staff3, DriverPg))

	items, result := Select(db, &staff)
	chkResult(t, result)

	staffs := items.([]Staff)

	if len(staffs) == 0 {
		t.Fatal("found 0 record")
	}

	for _, s := range staffs {
		t.Log(s.Id)
		if s.Name == "" {
			t.Fatal("Empty name")
		}
		if s.Id > 0 {
			var res = Delete(db, &s)
			if res.Error != nil {
				t.Fatal(res.Error)
			}
		}
	}
	_ = staffs
}
