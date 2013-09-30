package sqlutils

import "fmt"
import "strings"

func BuildWhereClauseWithAndOp(argMap map[string]interface{}, holderType int) (string, []interface{}) {
	sql, args := BuildWhereInnerClause(argMap, "AND", holderType)
	return " WHERE " + sql, args
}

func BuildWhereClauseWithOrOp(argMap map[string]interface{}, holderType int) (string, []interface{}) {
	sql, args := BuildWhereInnerClause(argMap, "OR", holderType)
	return " WHERE " + sql, args
}

func BuildWhereClauseWithOp(argMap map[string]interface{}, op string, holderType int) (string, []interface{}) {
	sql, args := BuildWhereInnerClause(argMap, op, holderType)
	return " WHERE " + sql, args
}

func BuildWhereInnerClause(argMap map[string]interface{}, op string, holderType int) (string, []interface{}) {
	var fields []string
	var args []interface{}
	var i int = 1
	for name, val := range argMap {
		if holderType == QMARK_HOLDER {
			fields = append(fields, fmt.Sprintf("%s = ?", name))
		} else {
			fields = append(fields, fmt.Sprintf("%s = $%d", name, i))
		}
		args = append(args, val)
		i++
	}
	return strings.Join(fields, fmt.Sprintf(" %s ", op)), args
}
