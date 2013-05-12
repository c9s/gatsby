package gatsby

import "database/sql"

/*
The record object, which is a struct pointer.
*/
type PtrRecord interface{}

type PtrTargetRecord interface {
	SetTarget(PtrRecord)
}

type BaseRecord struct {
	Txn    *sql.Tx
	Target PtrRecord
}

func NewRecord(o PtrTargetRecord) interface{} {
	o.SetTarget(o)
	return o
}

func (self *BaseRecord) SetTarget(target PtrRecord) {
	self.Target = target
}

func (self *BaseRecord) SetTxn(txn *sql.Tx) {
	self.Txn = txn
}

func (self *BaseRecord) GetTxn() *sql.Tx {
	return self.Txn
}

func (self *BaseRecord) Begin() (*sql.Tx, error) {
	var tx *sql.Tx
	var err error

	if self == nil {
		panic("You need to initialize the BaseRecord struct")
	}

	if self.Txn != nil {
		panic("You have already began a transaction.")
	}

	tx, err = conn.Begin()
	if err != nil {
		return nil, err
	}

	if tx == nil {
		panic("Empty transaction object, check driver support?")
	}

	self.Txn = tx
	return tx, nil
}

func (self *BaseRecord) Rollback() error {
	if self.Txn != nil {
		var err = self.Txn.Rollback()
		// free the transaction object
		self.Txn = nil
		return err
	}
	return nil
}

func (self *BaseRecord) Commit() error {
	if self.Txn != nil {
		var err = self.Txn.Commit()
		// free the transaction object
		self.Txn = nil
		return err
	}
	return nil
}

func (self *BaseRecord) Create() *Result {
	if self.Txn != nil {
		return Create(self.Txn, self.Target, driverType)
	}
	return Create(conn, self.Target, driverType)
}

func (self *BaseRecord) CreateWithInstance(o PtrRecord) *Result {
	if self.Txn != nil {
		return Create(self.Txn, o, driverType)
	}
	return Create(conn, o, driverType)
}

func (self *BaseRecord) Update() *Result {
	if self.Txn != nil {
		return Update(self.Txn, self.Target)
	}
	return Update(conn, self.Target)
}

func (self *BaseRecord) Load(id int64) *Result {
	return Load(conn, self.Target, id)
}

/*
Load and fill current record with customzied query (after the select clause)

	record.LoadWith(`WHERE name = $1`, name)

*/
func (self *BaseRecord) LoadWith(postQuery string, args ...interface{}) *Result {
	return LoadWith(conn, self.Target, postQuery, args...)
}

func (self *BaseRecord) LoadByCols(cols WhereMap) *Result {
	return LoadByCols(conn, self.Target, cols)
}

func (self *BaseRecord) Delete() *Result {
	// delete with transaction
	if self.Txn != nil {
		return Delete(self.Txn, self.Target)
	}
	return Delete(conn, self.Target)
}

func (self *BaseRecord) SelectByCols(cols WhereMap) (interface{}, *Result) {
	return SelectWhere(conn, self.Target, cols)
}

func (self *BaseRecord) DeleteWithInstance(o PtrRecord) *Result {
	// delete with transaction
	if self.Txn != nil {
		return Delete(self.Txn, o)
	}
	return Delete(conn, o)
}

func (self *BaseRecord) UpdateWithInstance(o PtrRecord) *Result {
	if self.Txn != nil {
		return Update(self.Txn, o)
	}
	return Update(conn, o)
}

func (self *BaseRecord) LoadWithInstance(o PtrRecord, id int64) *Result {
	return Load(conn, o, id)
}

func (self *BaseRecord) LoadByColsWithInstance(o PtrRecord, cols WhereMap) *Result {
	return LoadByCols(conn, o, cols)
}

func (self *BaseRecord) SelectByColsWithInstance(o PtrRecord, cols WhereMap) (interface{}, *Result) {
	return SelectWhere(conn, o, cols)
}
