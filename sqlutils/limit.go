package sqlutils

import "strconv"

/*
This function builds limit clause.
*/
func BuildLimitOffsetClause(limit int, offset int) string {
	return "LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa(offset)
}

func BuildLimitClause(limit int) string {
	return "LIMIT " + strconv.Itoa(limit)
}
