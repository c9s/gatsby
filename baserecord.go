package gatsby

import "gatsby/sqlutils"
import "database/sql"

type BaseRecord struct {
	Txn *sql.Tx
}

type EntityManager struct {
	BaseRecord
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

func (self *BaseRecord) Create(o PtrRecord) *Result {
	if self.Txn != nil {
		return Create(self.Txn, o, sqlutils.DriverPg)
	}
	return Create(conn, o, sqlutils.DriverPg)
}

func (self *BaseRecord) CreateWithInstance(o PtrRecord) *Result {
	if self.Txn != nil {
		return Create(self.Txn, o, sqlutils.DriverPg)
	}
	return Create(conn, o, sqlutils.DriverPg)
}

func (self *BaseRecord) Delete(o PtrRecord) *Result {
	// delete with transaction
	if self.Txn != nil {
		return Delete(self.Txn, o)
	}
	return Delete(conn, o)
}

func (self *BaseRecord) DeleteWithInstance(o PtrRecord) *Result {
	// delete with transaction
	if self.Txn != nil {
		return Delete(self.Txn, o)
	}
	return Delete(conn, o)
}

func (self *BaseRecord) Update(o PtrRecord) *Result {
	if self.Txn != nil {
		return Update(self.Txn, o)
	}
	return Update(conn, o)
}

func (self *BaseRecord) UpdateWithInstance(o PtrRecord) *Result {
	if self.Txn != nil {
		return Update(self.Txn, o)
	}
	return Update(conn, o)
}

func (self *BaseRecord) Load(o PtrRecord, id int64) *Result {
	return Load(conn, o, id)
}

func (self *BaseRecord) LoadWithInstance(o PtrRecord, id int64) *Result {
	return Load(conn, o, id)
}

func (self *BaseRecord) LoadByCols(o PtrRecord, cols WhereMap) *Result {
	return LoadByCols(conn, o, cols)
}

func (self *BaseRecord) LoadByColsWithInstance(o PtrRecord, cols WhereMap) *Result {
	return LoadByCols(conn, o, cols)
}

func (self *BaseRecord) SelectByColsWithInstance(o PtrRecord, cols WhereMap) (interface{}, *Result) {
	return SelectWhere(conn, o, cols)
}
