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
	GetPrimaryKeyValue() int64
	SetPrimaryKeyValue(int64)
}

type PrimaryKeyValue interface {
	GetPrimaryKeyValue() int64
}

type PrimaryKeyColumnName interface {
	GetPrimaryKeyColumnName() string
}

func SetPrimaryKeyValue(val interface{}, keyValue int64) bool {
	if idx := FindPrimaryKeyIdx(val); idx != -1 {
		t := reflect.ValueOf(val).Elem()
		t.Field(idx).SetInt(keyValue)
		return true
	}
	return false
}

func FindPrimaryKeyIdx(val interface{}) int {
	var cacheKey string = GetTypeName(val)
	if cache, ok := primaryKeyIdxCache[cacheKey]; ok {
		return cache
	}
	var t = reflect.ValueOf(val).Elem()
	var typeOfT = t.Type()
	var tag reflect.StructTag
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if HasColumnAttributeFromTag(&tag, "primary") {
			primaryKeyIdxCache[cacheKey] = i
			return i
		}
	}
	primaryKeyIdxCache[cacheKey] = -1
	return -1
}

// Find the primary key column and return the value of primary key.
// Return nil if primary key is not found.
func GetPrimaryKeyValue(val interface{}) *int64 {
	if o, ok := val.(PrimaryKeyValue); ok {
		i64 := o.GetPrimaryKeyValue()
		return &i64
	}

	if idx := FindPrimaryKeyIdx(val); idx != -1 {
		t := reflect.ValueOf(val).Elem()
		if val, ok := t.Field(idx).Interface().(int64); ok {
			return &val
		}
	}
	return nil
}

// Return the primary key column name, return nil if not found.
func GetPrimaryKeyColumnName(val interface{}) *string {
	var columnName *string
	var cacheKey string = GetTypeName(val)
	if cache, ok := primaryKeyColumnCache[cacheKey]; ok {
		return &cache
	}

	if o, ok := val.(PrimaryKeyColumnName); ok {
		name := o.GetPrimaryKeyColumnName()
		primaryKeyColumnCache[cacheKey] = name
		return &name
	} else {
		var tag reflect.StructTag
		var t = reflect.ValueOf(val).Elem()
		var typeOfT = t.Type()
		if idx := FindPrimaryKeyIdx(val); idx != -1 {
			tag = typeOfT.Field(idx).Tag
			if columnName = GetColumnNameFromTag(&tag); columnName != nil {
				primaryKeyColumnCache[cacheKey] = *columnName
				return columnName
			}
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
	var cacheKey string = GetTypeName(val)
	if cache, ok := columnNameClauseWithAliasCache[cacheKey]; ok {
		return cache
	}

	var t = reflect.ValueOf(val).Elem()
	var typeOfT = t.Type()
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
	columnNameClauseWithAliasCache[cacheKey] = sql
	return sql
}

func ReflectColumnNamesClause(val interface{}) string {

	var cacheKey string = GetTypeName(val)
	if cache, ok := columnNameClauseCache[cacheKey]; ok {
		return cache
	}

	var t = reflect.ValueOf(val).Elem()
	var typeOfT = t.Type()
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
	columnNameClauseCache[cacheKey] = sql
	return sql
}

// Iterate struct names and return a slice that contains column names.
func ReflectColumnNames(val interface{}) []string {

	var cacheKey string = GetTypeName(val)
	if cache, ok := columnNameCache[cacheKey]; ok {
		return cache
	}

	var t = reflect.ValueOf(val).Elem()
	var typeOfT = t.Type()
	var columns []string
	var columnName *string
	var tag reflect.StructTag
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName != nil {
			columns = append(columns, *columnName)
		}
	}
	columnNameCache[cacheKey] = columns
	return columns
}
