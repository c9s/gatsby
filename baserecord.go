package gatsby
import "gatsby/sqlutils"
// import "database/sql"

type BaseRecord struct {

}

func (self *BaseRecord) CreateWithInstance(o interface{}) (*sqlutils.Result) {
	return sqlutils.Create(appHandle.DbHandle, o, sqlutils.DriverPg)
}

func (self *BaseRecord) DeleteWithInstance(o interface{}) (*sqlutils.Result) {
	return sqlutils.Delete(appHandle.DbHandle, o)
}

func (self *BaseRecord) UpdateWithInstance(o interface{}) (*sqlutils.Result) {
	return sqlutils.Update(appHandle.DbHandle, o)
}

func (self *BaseRecord) LoadWithInstance(o interface{}, id int64) (*sqlutils.Result) {
	return sqlutils.Load(appHandle.DbHandle, o, id)
}


