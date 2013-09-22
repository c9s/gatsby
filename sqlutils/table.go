package sqlutils

import "reflect"
import "github.com/c9s/inflect"

// (struct pointer)
func GetTypeName(val interface{}) string {
	return reflect.ValueOf(val).Elem().Type().Name()
}

// (struct pointer)
func GetTableName(val interface{}) string {
	typeName := reflect.ValueOf(val).Elem().Type().Name()
	return GetTableNameFromTypeName(typeName)
}

// Convert type name to table name with underscore and plurize.
func GetTableNameFromTypeName(typeName string) string {
	if cache, ok := tableNameCache[typeName]; ok {
		return cache
	}
	tableNameCache[typeName] = inflect.Tableize(typeName)
	return tableNameCache[typeName]
}
