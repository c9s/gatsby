package sqlutils

import (
	"database/sql"
	"reflect"
)

func CreateReflectValuesFromTypes(types []interface{}) ([]interface{}, []reflect.Value) {
	var values []interface{}
	var reflectValues []reflect.Value

	for i := 0; i < len(types); i++ {
		var t = reflect.Indirect(reflect.ValueOf(types[i]))
		// var t = reflect.ValueOf( types[i] )
		var typeOfT = t.Type()
		// create val depends on types
		var value = reflect.New(typeOfT)
		reflectValues = append(reflectValues, value)
		values = append(values, value.Interface())
	}
	return values, reflectValues
}

func CreateMapsFromRows(rows *sql.Rows, types ...interface{}) ([]map[string]interface{}, error) {
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
		var result = map[string]interface{}{}
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

func CreateMapFromRows(rows *sql.Rows, types ...interface{}) (map[string]interface{}, error) {
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// create interface
	var values []interface{}
	var reflectValues []reflect.Value
	var result = map[string]interface{}{}

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
