package sqlutils

import "fmt"

/*
This function builds limit clause.
*/
func BuildLimitOffsetClause(limit int64, offset int64) string {
	return fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
}

func BuildLimitClause(limit int64) string {
	return fmt.Sprintf(" LIMIT %d", limit)
}
