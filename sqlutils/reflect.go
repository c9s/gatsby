package sqlutils
import "fmt"
import "reflect"
import "strings"
import "database/sql"
import _ "github.com/bmizerany/pq"

var columnNameCache = map[string] []string {};

func GetColumnMap(val interface{}) (map[string] interface{}) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()

	// var structName string = typeOfT.String()
	var columns = map[string] interface{} {};

	for i := 0; i < t.NumField(); i++ {
		var columnName string
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var field reflect.Value = t.Field(i)
		tagString := tag.Get("json")
		if len(tagString) > 0 {
			columnName = strings.SplitN(tagString,",",1)[0]
		}
		if len(columnName) == 0 {
			columnName = strings.SplitN(tag.Get("field"),",",1)[0]
		}
		if len(columnName) > 0 {
			columns[ columnName ] = field.Interface()
		}
	}
	return columns
}

// Parse SQL columns from struct
func ParseColumnNames(val interface{}) ([]string) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()

	var structName string = typeOfT.String()
	if cache, ok := columnNameCache[structName] ; ok {
		return cache
	}

	columns := []string{}
	for i := 0; i < t.NumField(); i++ {
		var columnName string
		var tag reflect.StructTag = typeOfT.Field(i).Tag

		var field reflect.Value = t.Field(i)
		fmt.Printf("%d: %s %s %s = %v\n", i,
			typeOfT.Field(i).Name,
			tag.Get("json"),
			field.Type(),
			field.Interface())

		tagString := tag.Get("json")

		if len(tagString) > 0 {
			columnName = strings.SplitN(tagString,",",1)[0]
		}


		if len(columnName) == 0 {
			columnName = strings.SplitN(tag.Get("field"),",",1)[0]
		}

		// XXX: use inflector to convert field name with underscore, maybe
		// columnName = typeOfT.Field(i).Name
		if len(columnName) > 0 {
			columns = append(columns,columnName)
		}
	}
	columnNameCache[structName] = columns
	return columns
}

// Generate SQL columns string for selecting.
func BuildSelectColumnClause(val interface{}) (string) {
	columns := ParseColumnNames(val)
	return strings.Join(columns,",")
}

func FillFromRow(val interface{}, rows * sql.Rows) (error) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var args []interface{}
	var fieldNum []int

	for i := 0; i < t.NumField(); i++ {
		var columnName string
		var tag       reflect.StructTag = typeOfT.Field(i).Tag
		var field     reflect.Value = t.Field(i)
		var fieldType reflect.Type = field.Type()

		tagString := tag.Get("json")
		if len(tagString) > 0 {
			columnName = strings.SplitN(tagString,",",1)[0]
		}
		if len(columnName) == 0 {
			columnName = strings.SplitN(tag.Get("field"),",",1)[0]
		}
		if len(columnName) == 0 {
			continue
		}

		// args = append(args, field.Interface())
		// args = append(args, field.Addr())
		// args = append(args, field.Elem() )
		if fieldType.String() == "string" {
			args = append(args, new(sql.NullString) )
		} else if fieldType.String() == "int" {
			args = append(args, new(sql.NullInt64) )
		} else if fieldType.String() == "bool" {
			args = append(args, new(sql.NullBool))
		} else if fieldType.String() == "float" {
			args = append(args, new(sql.NullFloat64))
		} else {
			// not sure if this work
			args = append(args, reflect.New(fieldType).Elem().Interface() )
		}
		fieldNum = append(fieldNum,i)
	}

	err := rows.Scan(args...)
	if err != nil {
		return err
	}

	for i, arg := range args {
		var fieldIdx int = fieldNum[i]
		var val reflect.Value = t.Field(fieldIdx)
		var t reflect.Type = val.Type()
		var typeStr string = t.String()

		if ! val.CanSet() {
			panic("can not set value " + typeOfT.Field(fieldIdx).Name + " on " + t.Name() )
		}


		if typeStr == "string" {
			if arg.(*sql.NullString).Valid {
				val.SetString( arg.(*sql.NullString).String)
			}
		} else if typeStr == "int" {
			if arg.(*sql.NullInt64).Valid {
				val.SetInt( arg.(*sql.NullInt64).Int64 )
			}
		} else if typeStr == "bool" {
			if arg.(*sql.NullBool).Valid {
				val.SetBool( arg.(*sql.NullBool).Bool)
			}
		} else if typeStr == "float" {
			if arg.(*sql.NullFloat64).Valid {
				val.SetFloat( arg.(*sql.NullFloat64).Float64)
			}
		} else {
			panic("unsupported type" + t.String() )
		}
	}
	return err
	// var id int
	// var name string
	// var stafftype string
	// var phone sql.NullString
	// var gender sql.NullString
	// return rows.Scan(&id, &name,  &gender, &stafftype, &phone)
}


