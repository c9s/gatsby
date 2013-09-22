package sqlutils

var selectQueryCache = map[string]string{}

// import "errors"

// Build SQL columns string for selecting,
// this function returns "column1, column2, column3"
func BuildSelectColumnClauseFromStruct(val interface{}) string {
	return ReflectColumnNamesClause(val)
}

// Build SQL columns string for selecting, this function returns
// "alias.column1, alias.column2, alias.column3"
func BuildSelectColumnClauseFromStructWithAlias(val interface{}, alias string) string {
	return ReflectColumnNamesClauseWithAlias(val, alias)
}

// Given a struct object, return a "SELECT ... FROM {tableName}" SQL clause.
func BuildSelectClause(val interface{}) string {
	k := GetTypeName(val)
	if sql, ok := selectQueryCache[k]; ok {
		return sql
	} else {
		sql := "SELECT " + BuildSelectColumnClauseFromStruct(val) + " FROM " + GetTableName(val)
		selectQueryCache[k] = sql
		return sql
	}
}

// Given a struct object and an alias string,
// This function returns a "SELECT alias1.column1, alias1.column2 FROM tableName alias" clause
func BuildSelectClauseWithAlias(val interface{}, alias string) string {
	return "SELECT " + BuildSelectColumnClauseFromStructWithAlias(val, alias) +
		" FROM " + GetTableName(val) + " " + alias
}
