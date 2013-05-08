package gatsby

import "gatsby/sqlutils"
import "database/sql"

type BaseRecord struct {
	Tx *sql.Tx
}

func (self *BaseRecord) CreateWithInstance(o interface{}) *Result {
	return Create(conn, o, sqlutils.DriverPg)
}

func (self *BaseRecord) DeleteWithInstance(o interface{}) *Result {
	// delete with transaction
	if self.Tx != nil {
		return Delete(self.Tx, o)
	}
	return Delete(conn, o)
}

func (self *BaseRecord) UpdateWithInstance(o interface{}) *Result {
	return Update(conn, o)
}

func (self *BaseRecord) LoadWithInstance(o interface{}, id int64) *Result {
	return Load(conn, o, id)
}

func (self *BaseRecord) LoadByColsWithInstance(o interface{}, cols map[string]interface{}) *Result {
	return LoadByCols(conn, o, cols)
}
