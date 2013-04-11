package sqlutils
import "fmt"
import "strings"


func BuildWhereClauseFromMap(argMap map[string]interface{}) (string, []interface{}) {
	var fields []string
	var args   []interface{}
	var i int = 1
	for name, val := range argMap {
		fields = append(fields, fmt.Sprintf("%s = $%d", name, i) )
		args   = append(args, val)
		i++
	}
	return " WHERE " + strings.Join(fields, " AND "), args
}

