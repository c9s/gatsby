package sqlutils

import "reflect"

// cache maps
var columnNameCache = map[string][]string{}
var tableNameCache = map[string]string{}

// provide PrimaryKey interface for faster column name accessing
type PrimaryKey interface {
	GetPkId() int64
	SetPkId(int64)
}

func SetPrimaryKeyValue(val interface{}, keyValue int64) bool {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		if GetColumnNameFromTag(&tag) == nil {
			continue
		}
		var columnAttributes = GetColumnAttributesFromTag(&tag)
		if _, ok := columnAttributes["primary"]; ok {
			t.Field(i).SetInt(keyValue)
			return true
		}
	}
	return false
}

// Find the primary key column and return the value of primary key.
// Return nil if primary key is not found.
func GetPrimaryKeyValue(val interface{}) *int64 {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		if GetColumnNameFromTag(&tag) == nil {
			continue
		}
		var columnAttributes = GetColumnAttributesFromTag(&tag)
		if _, ok := columnAttributes["primary"]; ok {
			val := t.Field(i).Interface().(int64)
			return &val
		}
	}
	return nil
}

// Return the primary key column name, return nil if not found.
func GetPrimaryKeyColumnName(val interface{}) *string {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		var columnAttributes = GetColumnAttributesFromTag(&tag)
		if _, ok := columnAttributes["primary"]; ok {
			return columnName
		}
	}
	return nil
}

// Iterate structure fields and return the
// values with map[string] interface{}
func GetColumnValueMap(val interface{}) map[string]interface{} {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	// var structName string = typeOfT.String()
	var columns = map[string]interface{}{}

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		columns[*columnName] = t.Field(i).Interface()
	}
	return columns
}

// Iterate struct names and return a slice that contains column names.
func ReflectColumnNames(val interface{}) []string {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var structName string = typeOfT.String()
	if cache, ok := columnNameCache[structName]; ok {
		return cache
	}

	var columns []string
	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		columns = append(columns, *columnName)
	}
	columnNameCache[structName] = columns
	return columns
}
