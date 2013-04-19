package gatsby

import "gatsby/sqlutils"
import "database/sql"

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



/*
func Select(val interface{}) (interface{}, *sqlutils.Result) {
	return sqlutils.Select(conn, val)
}

func SelectWith(val interface{}, postSql string, args ...interface{}) (interface{}, *sqlutils.Result) {
	return sqlutils.SelectWith(conn, val, postSql, args...)
}

func SelectWhere(val interface{}, conds map[string]interface{}) (interface{}, *sqlutils.Result) {
	return sqlutils.SelectWhere(conn, val, conds)
}

func SelectFromQuery(val interface{}, sql string, args ...interface{}) (interface{}, *sqlutils.Result) {
	return sqlutils.SelectFromQuery(conn, val, sql, args...)
}
