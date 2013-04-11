package sqlutils
import "reflect"
import "errors"
import "database/sql"


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

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}

		if fieldType.String() == "string" {
			args = append(args, new(sql.NullString) )
		} else if fieldType.String() == "int" {
			args = append(args, new(sql.NullInt64) )
		} else if fieldType.String() == "bool" {
			args = append(args, new(sql.NullBool))
		} else if fieldType.String() == "float" {
			args = append(args, new(sql.NullFloat64))
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
			return errors.New("can not set value " + typeOfT.Field(fieldIdx).Name + " on " + t.Name() )
		}

		// if arg.(*sql.NullString) == *sql.NullString {
		if typeStr == "string" {
			if arg.(*sql.NullString).Valid {
				val.SetString( arg.(*sql.NullString).String)
			} else if isRequired {
				return errors.New("required field")
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
			return errors.New("unsupported type" + t.String() )
		}
	}
	return err
}
