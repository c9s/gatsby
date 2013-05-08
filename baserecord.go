package gatsby

import "gatsby/sqlutils"
import "database/sql"

type BaseRecord struct {
	Txn *sql.Tx
}

func (self *BaseRecord) Begin() (*sql.Tx, error) {
	var tx *sql.Tx
	var err error

	tx, err = conn.Begin()
	if err != nil {
		return nil, err
	}

	if tx == nil {
		panic("empty transaction object, check driver support?")
	}
	if self == nil {
		panic("you need to initialize the BaseRecord struct")
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

func (self *BaseRecord) CreateWithInstance(o interface{}) *Result {
	if self.Txn != nil {
		return Create(self.Txn, o, sqlutils.DriverPg)
	}
	return Create(conn, o, sqlutils.DriverPg)
}

func (self *BaseRecord) DeleteWithInstance(o interface{}) *Result {
	// delete with transaction
	if self.Txn != nil {
		return Delete(self.Txn, o)
	}
	return Delete(conn, o)
}

func (self *BaseRecord) UpdateWithInstance(o interface{}) *Result {
	if self.Txn != nil {
		return Update(self.Txn, o)
	}
	return Update(conn, o)
}

func (self *BaseRecord) LoadWithInstance(o interface{}, id int64) *Result {
	return Load(conn, o, id)
}

func (self *BaseRecord) LoadByColsWithInstance(o interface{}, cols map[string]interface{}) *Result {
	return LoadByCols(conn, o, cols)
}
