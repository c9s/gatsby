package gatsby
import "fmt"
import "reflect"
// import "strings"


// Parse SQL columns from struct
func ParseSQLColumns(val interface{}) ([]string) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()
	columns := []string{}
	for i := 0; i < t.NumField(); i++ {
		var columnName string
		var field reflect.Value = t.Field(i)
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		fmt.Printf("%d: %s %s %s = %v\n", i,
			typeOfT.Field(i).Name,
			tag.Get("json"),
			field.Type(),
			field.Interface())

		columnName = tag.Get("json")
		if len(columnName) == 0 {
			columnName = typeOfT.Field(i).Name
		}
		columns = append(columns,columnName)
	}
	return columns
}

// Generate SQL columns string for selecting.
func GenerateSQLColumns(columns []string) (string) {
	columns := gatsby.ParseSQLColumns( drshine.Staff{Id:3} )
	return strings.Join(columns,",")
}

