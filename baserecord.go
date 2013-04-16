package gatsby
import "gatsby/sqlutils"
import "database/sql"

type BaseRecord struct {

}

var conn *sql.DB
var driverType int

func SetupConnection(c *sql.DB, driverType int) {
	conn = c
	driverType = driverType
}

func (self *BaseRecord) CreateWithInstance(o interface{}) (*sqlutils.Result) {
	return sqlutils.Create(conn, o, sqlutils.DriverPg)
}

func (self *BaseRecord) DeleteWithInstance(o interface{}) (*sqlutils.Result) {
	return sqlutils.Delete(conn, o)
}

func (self *BaseRecord) UpdateWithInstance(o interface{}) (*sqlutils.Result) {
	return sqlutils.Update(conn, o)
}

func (self *BaseRecord) LoadWithInstance(o interface{}, id int64) (*sqlutils.Result) {
	return sqlutils.Load(conn, o, id)
}

func (self *BaseRecord) LoadByColsWithInstance(o interface{}, id int64) (*sqlutils.Result) {
	return sqlutils.Load(conn, o, id)
}





