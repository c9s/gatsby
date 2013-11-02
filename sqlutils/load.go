package sqlutils

var loadQueryCache = map[string]string{}

func BuildLoadClause(val interface{}, holderType int) string {
	k := GetTypeName(val)
	if cache, ok := loadQueryCache[k]; ok {
		return cache
	} else {
		var pName = GetPrimaryKeyColumnName(val)
		if pName == nil {
			panic("primary key is required.")
		}
		var sqlstring = BuildSelectClause(val)
		if holderType == QMARK_HOLDER {
			sqlstring += " WHERE " + *pName + " = ?" + BuildLimitClause(1)
		} else {
			sqlstring += " WHERE " + *pName + " = $1" + BuildLimitClause(1)
		}
		loadQueryCache[k] = sqlstring
		return sqlstring
	}
}
