package gatsby
import "gatsby/sqlutils"
// import "database/sql"

func Load(val interface{}, pkId int64) (*sqlutils.Result) {
	return sqlutils.Load(conn, val, pkId)
}

func LoadByCols(val interface{}, cols map[string]interface{}) (*sqlutils.Result) {
	return sqlutils.LoadByCols(conn, val, cols)
}

func Create(val interface{}, driver int) (*sqlutils.Result) {
	return sqlutils.Create(conn, val, driver)
}

func Update(val interface{}) (*sqlutils.Result) {
	return sqlutils.Update(conn, val)
}

func Delete(val interface{}) (*sqlutils.Result) {
	return sqlutils.Delete(conn, val)
}
