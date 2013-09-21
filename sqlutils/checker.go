package sqlutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// struct pointer
func CheckRequired(structVal interface{}) error {
	t := reflect.ValueOf(structVal).Elem()
	typeOfT := t.Type()

	var tag reflect.StructTag
	var tagStr string
	var p int

	var fieldValue reflect.Value
	var fieldType reflect.StructField
	var val interface{}
	for i := 0; i < t.NumField(); i++ {
		tag = typeOfT.Field(i).Tag
		tagStr = tag.Get("field")

		// var attributes map[string]bool = GetColumnAttributesFromTag(&tag)
		// if _, ok := attributes["required"]; ok {
		if p = strings.Index(tagStr, "required"); p != -1 {
			fieldValue = t.Field(i)
			fieldType = typeOfT.Field(i)
			val = fieldValue.Interface()

			switch t := val.(type) {
			default:
				fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
			case string:
				if t == "" {
					return errors.New(fieldType.Name + " field is required.")
				}
			case int, int16, int32, int64:
				if t == 0 {
					return errors.New(fieldType.Name + " field is required.")
				}
			}
		}
	}
	return nil
}
