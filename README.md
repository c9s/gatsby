Gatsby Database Toolkit For Go
==============================


## Query Builder

Build Select Query:

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

Build Insert Query:

```go
import "gatsby"
query := NewQuery("staffs")
query.Insert(map[string]interface{} {
    "name": "John",
})
sql := query.String()
```

Build Update Query:

```go
query := NewQuery("staffs")
query.Update(map[string]interface{} {
    "name": "John",
})
query.WhereFromMap(map[string]interface{} {
    "id": 3,
})
sql := query.String()
```

