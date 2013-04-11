Gatsby SQLUtils
=================

SQLUtils package helps you build SQL.


```go
import "github.com/c9s/go-sqlutils" sqlutils

type Staff struct {
	Id        int `json:"id"`
	Name      string `json:"name" field:",required"`
	Gender    string `json:"gender"`
	StaffType string `json:"staff_type"` // valid types: doctor, nurse, ...etc
	Phone     string `json:"phone"`
}
sql := sqlutils.BuildSelectClause(Staff{})
// sql = " SELECT id, name, gender, staff_type, phone FROM staffs"

sql, args := sqlutils.BuildWhereClauseWithAndOp(map[string]interface{} {
    "name": "John"
})
// sql = " WHERE name = $1 AND id = $2"
```




