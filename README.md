Gatsby Database Toolkit For Go
==============================

## Query Builder

Gatsby Query provides a general query object to build SQL for selecting, updating, deleting, inserting records.

Currently it supports 4 mode (CRUD) to build SQL.

To Build Select Query:

```go
import "gatsby"

staffs := gatsby.NewQuery("staffs")
staffs.Select("id", "name")
staffs.WhereFromMap( gatsby.ArgMap{
    "name": "John"
})
sql := staffs.String()
args := staffs.Args()
```

To Build Insert Query:

```go
import "gatsby"
query := gatsby.NewQuery("staffs")
query.Insert(map[string]interface{} {
    "name": "John",
})
sql := query.String()
```

To Build Update Query:

```go
query := gatsby.NewQuery("staffs")
query.Update(map[string]interface{} {
    "name": "John",
})
query.WhereFromMap(map[string]interface{} {
    "id": 3,
})
sql := query.String()
```

## SQLFragments

More flexible SQL Builder by fragments.

You can append query fragments then combine them into one SQL string by joining, and you can use the generated SQL in anywhere you
want to combine with your own SQL statements.

SQLFragments filters these question marks into placeholders with number format, for example, the first `?` will be `$1`
and the second `?` will be `$2`.

```go
import "gatsby"
frag := gatsby.NewFragment()
frag.AppendQuery("name = ?", "John")
frag.AppendQuery("phone = ?", "John")
sql := frag.Join("OR")         // generates name = $1 AND phone = $2
args := frag.Args()
```

## BaseRecord

The BaseRecord provides a general CRUD operations on a struct type.

To define your model with Gatsby BaseRecord:

```go
package app
import "gatsby"

type Staff struct {
	Id        int64     `json:"id" field:"id,primary,serial"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	CellPhone string    `json:"cell_phone"`
	gatsby.BaseRecord
}
```

Then you can do CRUD operations on the struct object:

```go
staff := gatsby.NewRecord(&Staff{}).(*Staff)

res := staff.Load(10)   // load the record where primary key = 10

if res.Error != nil {
    // handle error here
}
if res.IsEmpty { 
    // handle empty record here
}

staff.Name = "John"

res := staff.Update()

res := staff.Delete()   // delete the record where primary key = 10

res := staff.Create()    // create another record
if res.Error != nil {

}
```


WIP
----

- Support uint64 for primary key.
- Support more NullInt types.
- Move pq.NullTime to support time.Time conversion.


