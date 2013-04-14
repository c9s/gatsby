Gatsby SQLUtils
=================

Gatsby SQLUtils package helps you build SQL through the struct object.

Importing
---------

Import from GitHub:

```go
import sqlutils "github.com/c9s/go-sqlutils"
```


Defining Struct Tag
-------------------

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
```

PrimaryKey interface
--------------------

Implement the PrimaryKey interface

```go
func (self *Staff) GetPkId() int64 {
    return self.Id
}

func (self *Staff) SetPkId(id int64) {
    self.Id = id
}
```



Buliding Select SQL statement
-----------------------------

To build select clause depends on the struct fields:

```go
sql := sqlutils.BuildSelectClause(&Staff{})
```

Which returns:

```sql
SELECT id, name, gender, staff_type, phone FROM staffs
```


```go
sql := sqlutils.BuildSelectColumnClause(&Staff{})
```

Which returns:

    id, name, gender, staff_type, phone


Building Update SQL statement
------------------------------

```go
sql, args := sqlutils.BuildUpdateClause(&Staff{})
```

Which returns:

```sql
UPDATE staffs SET name = $1, phone = $2, cellphone = $3
```

```go
sql, args := sqlutils.BuildUpdateColumns(&Staff{})
```

Which returns:

    name = $1, phone = $2, cellphone = $3

Building Where SQL statement
-----------------------------

To build where clause from map:

```
sql, args := sqlutils.BuildWhereClauseWithAndOp(map[string]interface{} {
    "name": "John"
})
```
which outputs:

```sql
WHERE name = $1
```

Which returns:

```sql
WHERE name = $1
```

Creating Record Through Database Connection
-------------------------------------------

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
rows, err := sqlutils.Update(db,&staff)
```









