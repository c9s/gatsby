Gatsby SQLUtils
=================

Gatsby SQLUtils package helps you build SQL.


Usage
-----

Import from GitHub:

```go
import "github.com/c9s/go-sqlutils" sqlutils
```

Define your struct with json spec or with field tag:

```go

type Staff struct {
	Id        int64  `field:",primary,serial"`
	Name      string `field:",required"`
	Gender    string `field:"gender"`
	StaffType string `field:"staff_type"`
	Phone     string `field:"phone"`
}

// to use with json tag
type Staff struct {
	Id        int64  `json:"id" field:",primary,serial"`
	Name      string `json:"name" field:",required"`
	Gender    string `json:"gender"`
	StaffType string `json:"staff_type"`
	Phone     string `json:"phone"`
}


// Implement the PrimaryKey interface
func (self *Staff) GetPkId() int64 {
    return self.Id
}
```

To build select clause depends on the struct fields:

```go
sql := sqlutils.BuildSelectClause(Staff{})
// sql = " SELECT id, name, gender, staff_type, phone FROM staffs"
```

To build where clause from map:

```
sql, args := sqlutils.BuildWhereClauseWithAndOp(map[string]interface{} {
    "name": "John"
})
// sql = " WHERE name = $1 AND id = $2"
```

To create new record:

```go
staff := Staff{Name:"Mary"}
result := sqlutils.Create(db,&staff, sqlutils.DriverPg)

if result.Error {
    // handle error
}
if result.Id != 0 {
    // handle primary key id
}
```

To update struct object:

```go
staff.Name = "NewName"
rows, err := Update(db,&staff)
```









