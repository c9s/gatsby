Gatsby Database Toolkit For Go
==============================

```go
import "gatsby"

staffs := gatsby.NewModel("staffs")
staffs.Select("id", "name")
staffs.WhereFromMap( gatsby.ArgMap{
    "name": "John"
})
sql := staffs.String()
args := staffs.Args()
```


