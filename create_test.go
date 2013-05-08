package gatsby

import "testing"

func TestCreate(t *testing.T) {
	var db = openDB()
	staff := Staff{}
	staff.Name = "John"
	Create(db, &staff, DriverPg)
	Delete(db, &staff)
}

func TestCreateWithTransactionAndCommit(t *testing.T) {
	var db = openDB()
	SetupConnection(db, DriverPg)
	staff := Staff{BaseRecord: &BaseRecord{}}
	staff.Name = "Txn Test"
	tx, err := staff.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("transaction is nil")
	}

	// create with transaction
	res := staff.Create()
	if res.Error != nil {
		t.Fatal(res.Error)
	}

	err = staff.Commit()
	if err != nil {
		t.Fatal(err)
	}
	CloseConnection()
}
