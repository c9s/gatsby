package sqlutils

// import "fmt"
import "reflect"
import "strings"

var columnNameCache = map[string] []string {};

func GetColumnMap(val interface{}) (map[string] interface{}) {
	// t := reflect.ValueOf(val)
	// typeOfT := t.Type()
	// var structName string = typeOfT.String()
	var columns = map[string] interface{} {};

	return columns
}

// Parse SQL columns from struct
func ParseColumnNames(val interface{}) ([]string) {
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

		tagString := tag.Get("json")

		if len(tagString) > 0 {
			columnName = strings.SplitN(tagString,",",1)[0]
		}


		if len(columnName) == 0 {
			columnName = strings.SplitN(tag.Get("field"),",",1)[0]
		}

		// XXX: use inflector to convert field name with underscore, maybe
		// columnName = typeOfT.Field(i).Name
		if len(columnName) > 0 {
			columns = append(columns,columnName)
		}
	}
	columnNameCache[structName] = columns
	return columns
}

// Generate SQL columns string for selecting.
func GenerateSelectColumns(val interface{}) (string) {
	columns := ParseColumnNames(val)
	return strings.Join(columns,",")
}


