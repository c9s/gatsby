package gatsby
import "gatsby/sqlutils"
import "database/sql"

type BaseRecord struct {

}

var conn *sql.DB

func SetupConnection(c *sql.DB) {
	conn = c
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



