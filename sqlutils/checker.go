package sqlutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func CheckRequired(structVal interface{}) error {
	if structVal == nil {
		return errors.New("Given struct value is nil.")
	}

	var t = reflect.ValueOf(structVal).Elem()
	var typeOfT = t.Type()
	var tag reflect.StructTag
	var tagStr string
	var p int

	var fieldType reflect.StructField
	var val interface{}
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		if tagStr = tag.Get("field"); len(tagStr) == 0 {
			continue
		}

		// var attributes map[string]bool = GetColumnAttributesFromTag(&tag)
		// if _, ok := attributes["required"]; ok {
		if p = strings.Index(tagStr, "required"); p != -1 && p > 0 {
			val = t.Field(i).Interface()

			switch t := val.(type) {
			default:
				fmt.Printf("unsuppported type for required attribute: %T\n", t) // %T prints whatever type t has
			case string:
				if t == "" {
					fieldType = typeOfT.Field(i)
					return errors.New(fieldType.Name + " field is required.")
				}
			case int, int8, int16, int32, int64:
				if t == 0 {
					fieldType = typeOfT.Field(i)
					return errors.New(fieldType.Name + " field is required.")
				}
			}
		}
	}
	return nil
}
