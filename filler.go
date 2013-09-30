package gatsby

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c9s/gatsby/sqlutils"
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
	var fieldValue reflect.Value
	var rval interface{}

	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag

		if sqlutils.GetColumnNameFromTag(&tag) == nil {
			continue
		}

		fieldValue = t.Field(i)

		rval = fieldValue.Interface()

		switch rval.(type) {
		case string:
			args = append(args, new(sql.NullString))
		case int64, int32, int16, int8, int:
			args = append(args, new(sql.NullInt64))
		case bool:
			args = append(args, new(sql.NullBool))
		case float64:
			args = append(args, new(sql.NullFloat64))
		case *time.Time:
			args = append(args, new(pq.NullTime))
		default:
			// XXX: Not sure if this work
			var fieldType reflect.Type = fieldValue.Type()
			args = append(args, reflect.New(fieldType).Elem().Interface())
		}

		fieldNums = append(fieldNums, i)
	}

	if err = rows.Scan(args...); err != nil {
		return err
	}

	var fieldIdx int

	for i, arg := range args {
		fieldIdx = fieldNums[i]
		// tag = typeOfT.Field(fieldIdx).Tag
		// isRequired := sqlutils.HasColumnAttributeFromTag(&tag, "required")

		fieldValue = t.Field(fieldIdx)
		// var rval = fieldValue.Interface()

		if !fieldValue.CanSet() {
			var valueType reflect.Type = fieldValue.Type()
			return errors.New("Can not set value " + typeOfT.Field(fieldIdx).Name + " on " + valueType.Name())
		}

		switch argVal := arg.(type) {
		case *sql.NullString:
			if argVal.Valid {
				fieldValue.SetString(argVal.String)
			}
		case *sql.NullInt64:
			if argVal.Valid {
				fieldValue.SetInt(argVal.Int64)
			}
		case *sql.NullBool:
			if argVal.Valid {
				fieldValue.SetBool(argVal.Bool)
			}
		case *sql.NullFloat64:
			if argVal.Valid {
				fieldValue.SetFloat(argVal.Float64)
			}
		case *pq.NullTime:
			if argVal.Valid {
				fieldValue.Set(reflect.ValueOf(&argVal.Time))
			}
		default:
			return fmt.Errorf("unsupported type %T", argVal)
		}
	}
	return err
}

func CreateMapsFromRows(rows *sql.Rows, types ...interface{}) (RecordMapList, error) {
	var err error

	columnNames, err := rows.Columns()
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
		if err = rows.Scan(values...); err != nil {
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
	var err error
	var values []interface{}
	var reflectValues []reflect.Value
	var result = RecordMap{}

	values, reflectValues = sqlutils.CreateReflectValuesFromTypes(types)
	if err = rows.Scan(values...); err != nil {
		return nil, err
	}

	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for i, n := range columnNames {
		result[n] = reflectValues[i].Elem().Interface()
	}
	return result, nil
}
