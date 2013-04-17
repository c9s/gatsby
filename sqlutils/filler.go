package sqlutils
import "reflect"
import "errors"
import "database/sql"
// import "time"
import "github.com/c9s/pq"

func CreateReflectValuesFromTypes(types []interface{}) ([]interface{}, []reflect.Value ) {
	var values []interface{}
	var reflectValues []reflect.Value

	for i := 0 ; i < len(types) ; i++ {
		var t = reflect.Indirect( reflect.ValueOf( types[i] ) )
		// var t = reflect.ValueOf( types[i] )
		var typeOfT = t.Type()
		// create val depends on types
		var value = reflect.New(typeOfT)
		reflectValues = append(reflectValues, value)
		values = append(values, value.Interface() )
	}
	return values, reflectValues
}


func CreateMapsFromRows(rows * sql.Rows, types ...interface{}) ([]map[string]interface{}, error) {
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// create interface
	var values []interface{}
	var reflectValues []reflect.Value

	var results []map[string]interface{}
	values, reflectValues = CreateReflectValuesFromTypes(types)

	for rows.Next() {
		var result = map[string]interface{} {}
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		for i, name := range columnNames {
			result[name] = reflectValues[i].Elem().Interface()
		}
		results = append(results,result)
	}
	return results, nil
}

func CreateMapFromRows(rows * sql.Rows, types ...interface{}) (map[string]interface{}, error) {
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// create interface
	var values []interface{}
	var reflectValues []reflect.Value
	var result = map[string]interface{} {}

	values, reflectValues = CreateReflectValuesFromTypes(types)
	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}
	for i, n := range columnNames {
		result[n] = reflectValues[i].Elem().Interface()
	}
	return result, nil
}



// Fill the struct data from a result rows
// This function iterates the struct by reflection, and creates types from sql
// package for filling result.
func FillFromRow(val interface{}, rows * sql.Rows) (error) {
	t       := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var args      []interface{}
	var fieldNums []int
	var fieldAttrList []map[string] bool

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)
		var fieldType  reflect.Type = field.Type()

		if tag.Get("field") == "-" {
			continue
		}

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}

		var typeStr string = fieldType.String()

		if typeStr == "string" {
			args = append(args, new(sql.NullString) )
		} else if typeStr == "int" || typeStr == "int64" {
			args = append(args, new(sql.NullInt64) )
		} else if typeStr == "bool" {
			args = append(args, new(sql.NullBool))
		} else if typeStr == "float" || typeStr == "float64" {
			args = append(args, new(sql.NullFloat64))
		} else if typeStr == "*time.Time" {
			args = append(args, new(pq.NullTime))
		} else {
			// Not sure if this work
			args = append(args, reflect.New(fieldType).Elem().Interface() )
		}

		fieldAttrs := GetColumnAttributesFromTag(&tag)

		fieldNums = append(fieldNums,i)
		fieldAttrList = append(fieldAttrList, fieldAttrs)
	}

	err := rows.Scan(args...)
	if err != nil {
		return err
	}

	for i, arg := range args {
		var fieldIdx int = fieldNums[i]
		fieldAttrs := fieldAttrList[i]

		var isRequired = fieldAttrs["required"]
		var val reflect.Value = t.Field(fieldIdx)
		var t reflect.Type = val.Type()
		var typeStr string = t.String()

		if ! val.CanSet() {
			return errors.New("Can not set value " + typeOfT.Field(fieldIdx).Name + " on " + t.Name() )
		}

		// if arg.(*sql.NullString) == *sql.NullString {
		if typeStr == "string" {
			if arg.(*sql.NullString).Valid {
				val.SetString( arg.(*sql.NullString).String)
			} else if isRequired {
				return errors.New("required field")
			}
		} else if typeStr == "int" || typeStr == "int64" {
			if arg.(*sql.NullInt64).Valid {
				val.SetInt( arg.(*sql.NullInt64).Int64 )
			}
		} else if typeStr == "bool" {
			if arg.(*sql.NullBool).Valid {
				val.SetBool( arg.(*sql.NullBool).Bool)
			}
		} else if typeStr == "float" || typeStr == "float64" {
			if arg.(*sql.NullFloat64).Valid {
				val.SetFloat( arg.(*sql.NullFloat64).Float64)
			}
		} else if typeStr == "*time.Time" {
			if nullTimeVal, ok := arg.(*pq.NullTime) ; ok && nullTimeVal != nil {
				val.Set( reflect.ValueOf(&nullTimeVal.Time) )
			}
		} else {
			return errors.New("unsupported type " + t.String() )
		}
	}
	return err
}
