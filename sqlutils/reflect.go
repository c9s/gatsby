package sqlutils
// import "fmt"
import "reflect"
import "strings"
// import "github.com/c9s/inflect"
import _ "github.com/bmizerany/pq"

var columnNameCache = map[string] []string {};

func GetColumnNameFromTag(tag reflect.StructTag) (*string) {
	fieldTags := strings.Split(tag.Get("field"),",")
	if len(fieldTags[0]) > 0 {
		return &fieldTags[0]
	}
	jsonTags := strings.Split(tag.Get("json"),",")
	if len(jsonTags[0]) == 0 {
		return &jsonTags[0]
	}
	return nil
}

func GetColumnMap(val interface{}) (map[string] interface{}) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()

	// var structName string = typeOfT.String()
	var columns = map[string] interface{} {};

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var field reflect.Value = t.Field(i)

		columnName := GetColumnNameFromTag(tag)
		if columnName != nil {
			columns[ *columnName ] = field.Interface()
		}
	}
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
		var tag reflect.StructTag = typeOfT.Field(i).Tag

		/*
		// var field reflect.Value = t.Field(i)
		fmt.Printf("%d: %s %s %s = %v\n", i,
			typeOfT.Field(i).Name,
			tag.Get("json"),
			field.Type(),
			field.Interface())
		*/

		columnName := GetColumnNameFromTag(tag)
		if columnName != nil {
			columns = append(columns,*columnName)
		}
	}
	columnNameCache[structName] = columns
	return columns
}





