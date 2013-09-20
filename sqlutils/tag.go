package sqlutils

import (
	"reflect"
	"strings"
)

// Extract column name attribute from struct tag (the first element) of the 'field' tag or
// column name from 'json' tag.
func GetColumnNameFromTag(tag *reflect.StructTag) *string {

	if tagStr := tag.Get("field"); len(tagStr) != 0 {
		// ignore it if it starts with dash
		if tagStr[0:1] == "-" {
			return nil
		}
		if p := strings.Index(tagStr, ","); p != -1 && p > 1 {
			str := tagStr[:p]
			return &str
		}
	}

	jsonTagStr := tag.Get("json")
	if len(jsonTagStr) == 0 {
		return nil
	}
	if jsonTagStr[0:1] == "-" {
		return nil
	}
	if p := strings.Index(jsonTagStr, ","); p != -1 {
		if p > 1 {
			str := jsonTagStr[:p]
			return &str
		}
		return nil
	}
	return &jsonTagStr
}

func HasColumnAttributeFromTag(tag *reflect.StructTag, aName string) bool {
	tagStr := tag.Get("field")
	return strings.Index(tagStr, ","+aName) != -1
}

// Extract attributes from "field" tag.
// Current supported attributes: "required","primary","serial"
func GetColumnAttributesFromTag(tag *reflect.StructTag) map[string]bool {
	fieldTags := strings.Split(tag.Get("field"), ",")
	attributes := map[string]bool{}
	for _, tag := range fieldTags[1:] {
		attributes[tag] = true
	}
	return attributes
}
