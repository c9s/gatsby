package gatsby

import (
	"database/sql"
	"errors"
	"fmt"
	"gatsby/sqlutils"
	"github.com/c9s/pq"
	"reflect"
	"time"
)

type RowScanner interface {
	Scan(dest ...interface{}) error
}

type RecordMap map[string]interface{}

type RecordMapList []RecordMap

// Fill the struct data from a result rows
// This function iterates the struct by reflection, and creates types from sql
// package for filling result.
func FillFromRows(val PtrRecord, rows RowScanner) error {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	var err error
	var args []interface{}
	var tag reflect.StructTag
	var fieldNums []int

	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag

		if sqlutils.GetColumnNameFromTag(&tag) == nil {
			continue
		}

		var field reflect.Value = t.Field(i)
		var fieldValue = field.Interface()

		switch fieldValue.(type) {
		case string:
			args = append(args, new(sql.NullString))
		case int64, int32, int16, int:
			args = append(args, new(sql.NullInt64))
		case bool:
			args = append(args, new(sql.NullBool))
		case float64:
			args = append(args, new(sql.NullFloat64))
		case *time.Time:
			args = append(args, new(pq.NullTime))
		default:
			// XXX: Not sure if this work
			var fieldType reflect.Type = field.Type()
			args = append(args, reflect.New(fieldType).Elem().Interface())
		}

		fieldNums = append(fieldNums, i)
	}

	err = rows.Scan(args...)
	if err != nil {
		return err
	}

	for i, arg := range args {
		var fieldIdx int = fieldNums[i]
		// tag = typeOfT.Field(fieldIdx).Tag
		// isRequired := sqlutils.HasColumnAttributeFromTag(&tag, "required")

		var fieldValue reflect.Value = t.Field(fieldIdx)
		var val = fieldValue.Interface()

		if !fieldValue.CanSet() {
			var valueType reflect.Type = fieldValue.Type()
			return errors.New("Can not set value " + typeOfT.Field(fieldIdx).Name + " on " + valueType.Name())
		}

		switch realVal := val.(type) {
		case string:
			if arg.(*sql.NullString).Valid {
				fieldValue.SetString(arg.(*sql.NullString).String)
			}
		case int, int16, int32, int64:
			if arg.(*sql.NullInt64).Valid {
				fieldValue.SetInt(arg.(*sql.NullInt64).Int64)
			}
		case bool:
			if arg.(*sql.NullBool).Valid {
				fieldValue.SetBool(arg.(*sql.NullBool).Bool)
			}
		case float64:
			if arg.(*sql.NullFloat64).Valid {
				fieldValue.SetFloat(arg.(*sql.NullFloat64).Float64)
			}
		case *time.Time:
			if nullTimeVal, ok := arg.(*pq.NullTime); ok && nullTimeVal != nil {
				fieldValue.Set(reflect.ValueOf(&nullTimeVal.Time))
			}
		default:
			return fmt.Errorf("unsupported type %T", realVal)
		}
	}
	return err
}

func CreateMapsFromRows(rows *sql.Rows, types ...interface{}) (RecordMapList, error) {
	var columnNames, err = rows.Columns()
	if err != nil {
		return nil, err
	}

	// create interface
	var values []interface{}
	var reflectValues []reflect.Value

	var results RecordMapList
	values, reflectValues = sqlutils.CreateReflectValuesFromTypes(types)

	for rows.Next() {
		var result = RecordMap{}
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		for i, name := range columnNames {
			result[name] = reflectValues[i].Elem().Interface()
		}
		results = append(results, result)
	}
	return results, nil
}

func CreateMapFromRows(rows *sql.Rows, types ...interface{}) (RecordMap, error) {
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// create interface
	var values []interface{}
	var reflectValues []reflect.Value
	var result = RecordMap{}

	values, reflectValues = sqlutils.CreateReflectValuesFromTypes(types)
	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}
	for i, n := range columnNames {
		result[n] = reflectValues[i].Elem().Interface()
	}
	return result, nil
}
