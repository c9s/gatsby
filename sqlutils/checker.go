package sqlutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// struct pointer
func CheckRequired(val interface{}) error {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag reflect.StructTag = typeOfT.Field(i).Tag
		var tagStr = tag.Get("field")

		// var attributes map[string]bool = GetColumnAttributesFromTag(&tag)
		// if _, ok := attributes["required"]; ok {
		if p := strings.Index(tagStr, "required"); p != -1 {
			var fieldValue reflect.Value = t.Field(i)
			var fieldType = typeOfT.Field(i)
			var val = fieldValue.Interface()

			switch t := val.(type) {
			default:
				fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
			case string:
				if t == "" {
					return errors.New(fieldType.Name + " field is required.")
				}
			case int:
				if t == 0 {
					return errors.New(fieldType.Name + " field is required.")
				}
			}
		}
	}
	return nil
}
