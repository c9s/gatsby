package sqlutils

var loadQueryCache = map[string]string{}

func BuildLoadClause(val interface{}) string {
	n := GetTableName(val)
	if cache, ok := loadQueryCache[n]; ok {
		return cache
	} else {
		var pName = GetPrimaryKeyColumnName(val)
		if pName == nil {
			panic("primary key is required.")
		}
		var sqlstring = BuildSelectClause(val) + " WHERE " + *pName + " = $1" + BuildLimitClause(1)
		loadQueryCache[n] = sqlstring
		return sqlstring
	}
}
