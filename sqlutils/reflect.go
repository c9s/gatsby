package sqlutils

// import "fmt"
import "reflect"
import "strings"

var columnNameCache = map[string] []string {};

// Parse SQL columns from struct
func ParseColumns(val interface{}) ([]string) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()

	var structName string = typeOfT.String()

	if cache, ok := columnNameCache[structName] ; ok {
		return cache
	}

	columns := []string{}
	for i := 0; i < t.NumField(); i++ {
		var columnName string
		var tag reflect.StructTag = typeOfT.Field(i).Tag

		// var field reflect.Value = t.Field(i)
		/*
		fmt.Printf("%d: %s %s %s = %v\n", i,
			typeOfT.Field(i).Name,
			tag.Get("json"),
			field.Type(),
			field.Interface())
		*/

		columnName = tag.Get("json")
		if len(columnName) == 0 {
			columnName = typeOfT.Field(i).Name
		}
		columns = append(columns,columnName)
	}
	columnNameCache[structName] = columns
	return columns
}

// Generate SQL columns string for selecting.
func GenerateSelectColumns(val interface{}) (string) {
	columns := ParseColumns(val)
	return strings.Join(columns,",")
}


