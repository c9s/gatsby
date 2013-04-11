package sqlutils
import "fmt"
import "strings"


func BuildWhereClauseWithAndOp(argMap map[string]interface{}) (string, []interface{}) {
	sql, args := BuildWhereInnerClause(argMap, "AND")
	return " WHERE " + sql, args
}

func BuildWhereClauseWithOrOp(argMap map[string]interface{}) (string, []interface{}) {
	sql, args := BuildWhereInnerClause(argMap, "OR")
	return " WHERE " + sql, args
}

func BuildWhereClauseWithOp(argMap map[string]interface{}, op string) (string, []interface{}) {
	sql, args := BuildWhereInnerClause(argMap, op)
	return " WHERE " + sql, args
}

func BuildWhereInnerClause(argMap map[string]interface{}, op string) (string, []interface{}) {
	var fields []string
	var args   []interface{}
	var i int = 1
	for name, val := range argMap {
		fields = append(fields, fmt.Sprintf("%s = $%d", name, i) )
		args   = append(args, val)
		i++
	}
	return strings.Join(fields, fmt.Sprintf(" %s ", op) ), args
}
