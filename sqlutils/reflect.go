package sqlutils

import "reflect"

// cache maps
var columnNameCache = map[string][]string{}
var columnNameClauseCache = map[string]string{}
var columnNameClauseWithAliasCache = map[string]string{}
var tableNameCache = map[string]string{}
var primaryKeyColumnCache = map[string]string{}
var primaryKeyIdxCache = map[string]int{}

// provide PrimaryKey interface for faster column name accessing
type PrimaryKey interface {
	GetPkId() int64
	SetPkId(int64)
}

func SetPrimaryKeyValue(val interface{}, keyValue int64) bool {
	if idx := GetPrimaryKeyIdx(val); idx != -1 {
		t := reflect.ValueOf(val).Elem()
		t.Field(idx).SetInt(keyValue)
		return true
	}
	return false
}

func GetPrimaryKeyIdx(val interface{}) int {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()
	var name string = typeOfT.String()
	if cache, ok := primaryKeyIdxCache[name]; ok {
		return cache
	}
	var tag reflect.StructTag
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if HasColumnAttributeFromTag(&tag, "primary") {
			primaryKeyIdxCache[name] = i
			return i
		}
	}
	primaryKeyIdxCache[name] = -1
	return -1
}

// Find the primary key column and return the value of primary key.
// Return nil if primary key is not found.
func GetPrimaryKeyValue(val interface{}) *int64 {
	if idx := GetPrimaryKeyIdx(val); idx != -1 {
		t := reflect.ValueOf(val).Elem()
		if val, ok := t.Field(idx).Interface().(int64); ok {
			return &val
		}
	}
	return nil
}

// Return the primary key column name, return nil if not found.
func GetPrimaryKeyColumnName(val interface{}) *string {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var columnName *string
	var structName string = typeOfT.String()
	var tag reflect.StructTag

	if cache, ok := primaryKeyColumnCache[structName]; ok {
		return &cache
	}

	if idx := GetPrimaryKeyIdx(val); idx != -1 {
		tag = typeOfT.Field(idx).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName != nil {
			primaryKeyColumnCache[structName] = *columnName
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
	var tag reflect.StructTag
	var columnName *string
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName != nil {
			columns[*columnName] = t.Field(i).Interface()
		}
	}
	return columns
}

func ReflectColumnNamesClauseWithAlias(val interface{}, alias string) string {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var structName string = typeOfT.String()
	if cache, ok := columnNameClauseWithAliasCache[structName]; ok {
		return cache
	}

	var sql string = ""
	var columnName *string
	var tag reflect.StructTag
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName != nil {
			sql += alias + "." + *columnName + ", "
		}
	}

	sql = sql[:len(sql)-2]
	columnNameClauseWithAliasCache[structName] = sql
	return sql
}

func ReflectColumnNamesClause(val interface{}) string {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var structName string = typeOfT.String()
	if cache, ok := columnNameClauseCache[structName]; ok {
		return cache
	}

	var sql string = ""
	var columnName *string
	var tag reflect.StructTag
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName != nil {
			sql += *columnName + ", "
		}
	}

	sql = sql[:len(sql)-2]
	columnNameClauseCache[structName] = sql
	return sql
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
	var columnName *string
	var tag reflect.StructTag
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName != nil {
			columns = append(columns, *columnName)
		}
	}
	columnNameCache[structName] = columns
	return columns
}
