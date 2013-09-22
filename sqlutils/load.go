package sqlutils

var loadQueryCache = map[string]string{}

func BuildLoadClause(val interface{}) string {
	k := GetTypeName(val)
	if cache, ok := loadQueryCache[k]; ok {
		return cache
	} else {
		var pName = GetPrimaryKeyColumnName(val)
		if pName == nil {
			panic("primary key is required.")
		}
		var sqlstring = BuildSelectClause(val) + " WHERE " + *pName + " = $1" + BuildLimitClause(1)
		loadQueryCache[k] = sqlstring
		return sqlstring
	}
}
