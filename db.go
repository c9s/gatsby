package gatsby
import "database/sql"

const (
	DriverPg = iota
	DriverMysql
	DriverSqlite
)

var conn *sql.DB
var driverType int

func SetupConnection(c *sql.DB, driverType int) {
	conn = c
	driverType = driverType
}

func GetConnection() *sql.DB {
	return conn
}


type ConnectionHandle struct {
	conn *sql.DB
	driverType int
}

func (self * ConnectionHandle) Load(val interface{}, pkId int64) *Result {
	return Load(self.conn, val, pkId)
}

func (self * ConnectionHandle) LoadByCols(val interface{}, cols map[string]interface{}) *Result {
	return LoadByCols(self.conn, val, cols)
}

func (self * ConnectionHandle) Create(val interface{}, driver int) *Result {
	return Create(self.conn, val, driver)
}


func (self * ConnectionHandle) Update(val interface{}) *Result {
	return Update(self.conn, val)
}

func (self * ConnectionHandle) Delete(val interface{}) *Result {
	return Delete(self.conn, val)
}




