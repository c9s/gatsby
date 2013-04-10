package sqlutils
import "strings"
import "reflect"
import "github.com/c9s/inflect"
import "database/sql"

// Generate SQL columns string for selecting.
func BuildSelectColumnClause(val interface{}) (string) {
	columns := ParseColumnNames(val)
	return strings.Join(columns,",")
}

func BuildSelectClause(val interface{}) (string) {
	// get table name
	// inflect.Underscore()
	t := reflect.ValueOf(val)
	typeOfT := t.Type()
	tableName := inflect.Tableize(typeOfT.Name())
	return "SELECT " + BuildSelectColumnClause(val) + " FROM " + tableName;
}

func PrepareAndQuery(db *sql.DB, sql string, args ...interface{}) (*sql.Rows,error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

