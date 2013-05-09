package sqlutils

import "errors"
import "reflect"

// struct pointer
func CheckRequired(val interface{}) error {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var fieldValue reflect.Value = t.Field(i)
		var fieldType = typeOfT.Field(i)
		var attributes map[string]bool = GetColumnAttributesFromTag(&tag)
		if _, ok := attributes["required"]; ok {
			// check the column value
			if fieldValue.Type().Name() == "string" && fieldValue.Interface().(string) == "" {
				return errors.New(fieldType.Name + " field is required.")
			} else if fieldValue.Type().Name() == "int" && fieldValue.Interface().(int) == 0 {
				return errors.New(fieldType.Name + " field is required.")
			}
		}
	}
	return nil
}
