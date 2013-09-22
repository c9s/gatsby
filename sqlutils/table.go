package sqlutils

import "reflect"
import "github.com/c9s/inflect"

// (struct pointer)
func GetTypeName(val interface{}) string {
	var v = reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		return v.Elem().Type().Name()
	}
	return v.Type().Name()
}

// (struct pointer)
func GetTableName(val interface{}) string {
	return GetTableNameFromTypeName(GetTypeName(val))
}

// Convert type name to table name with underscore and plurize.
func GetTableNameFromTypeName(typeName string) string {
	if cache, ok := tableNameCache[typeName]; ok {
		return cache
	}
	tableNameCache[typeName] = inflect.Tableize(typeName)
	return tableNameCache[typeName]
}
