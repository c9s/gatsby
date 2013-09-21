package sqlutils

import (
	"reflect"
)

func CreateReflectValuesFromTypes(types []interface{}) ([]interface{}, []reflect.Value) {
	var values []interface{}
	var reflectValues []reflect.Value
	var t reflect.Value
	var typeOfT reflect.Type

	for i := 0; i < len(types); i++ {
		t = reflect.Indirect(reflect.ValueOf(types[i]))
		// var t = reflect.ValueOf( types[i] )
		typeOfT = t.Type()

		// create val depends on types
		var value = reflect.New(typeOfT)
		reflectValues = append(reflectValues, value)
		values = append(values, value.Interface())
	}
	return values, reflectValues
}
