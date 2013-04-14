package sqlutils
import "reflect"
import "strings"
// import "fmt"
import "github.com/c9s/inflect"
import _ "github.com/bmizerany/pq"

var columnNameCache = map[string] []string {};
var tableNameCache = map[string] string {};

// provide PrimaryKey interface for faster column name accessing 
type PrimaryKey interface {
	GetPkId() int64
	SetPkId(int64)
}

func GetTableName(val interface{}) (string) {
	typeName := reflect.ValueOf(val).Elem().Type().Name()
	return GetTableNameFromTypeName(typeName)
}



// Find the primary key column and return the value of primary key.
// Return nil if primary key is not found.
func GetPrimaryKeyValue(val interface{}) *int64 {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var field reflect.Value = t.Field(i)
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		var columnAttributes = GetColumnAttributesFromTag(&tag)
		if _, ok := columnAttributes["primary"] ; ok {
			val := field.Interface().(int64)
			return &val
		}
	}
	return nil

}

// Return the primary key column name, return nil if not found.
func GetPrimaryKeyColumnName(val interface{}) (*string) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		var columnAttributes = GetColumnAttributesFromTag(&tag)
		if _, ok := columnAttributes["primary"] ; ok {
			return columnName
		}
	}
	return nil
}


// Convert type name to table name with underscore and plurize.
func GetTableNameFromTypeName(typeName string) (string) {
	if cache, ok := tableNameCache[ typeName ] ; ok {
		return cache
	}
	tableNameCache[typeName] = inflect.Tableize(typeName)
	return tableNameCache[typeName]
}


// Extract attributes from "field" tag.
// Current supported attributes: "required","primary","serial"
func GetColumnAttributesFromTag(tag *reflect.StructTag) (map[string]bool) {
	fieldTags := strings.Split(tag.Get("field"),",")
	attributes := map[string]bool {}
	for _, tag := range fieldTags[1:] {
		attributes[tag] = true
	}
	return attributes
}

// Extract column name attribute from struct tag (the first element) of the 'field' tag or 
// column name from 'json' tag.
func GetColumnNameFromTag(tag *reflect.StructTag) (*string) {
	fieldTags := strings.Split(tag.Get("field"),",")
	if len(fieldTags[0]) > 0 {
		return &fieldTags[0]
	}
	jsonTags := strings.Split(tag.Get("json"),",")
	if len(jsonTags[0]) > 0 {
		return &jsonTags[0]
	}
	return nil
}


// Iterate structure fields and return the 
// values with map[string] interface{}
func GetColumnValueMap(val interface{}) (map[string] interface{}) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	// var structName string = typeOfT.String()
	var columns = map[string] interface{} {};

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var field reflect.Value = t.Field(i)
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName != nil {
			columns[ *columnName ] = field.Interface()
		}
	}
	return columns
}

// Iterate struct names and return a slice that contains column names.
func ParseColumnNames(val interface{}) ([]string) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var structName string = typeOfT.String()
	if cache, ok := columnNameCache[structName] ; ok {
		return cache
	}

	var columns []string
	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName != nil {
			columns = append(columns, *columnName)
		}
	}
	columnNameCache[structName] = columns
	return columns
}





