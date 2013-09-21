package sqlutils

import (
	"reflect"
	"strings"
)

func IndexOfChar(s string, sep string) int {
	var l = len(s)
	var c = sep[0]
	for i := 0; i < l; i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// Extract column name attribute from struct tag (the first element) of the 'field' tag or
// column name from 'json' tag.
func GetColumnNameFromTag(tag *reflect.StructTag) *string {
	var p int
	var tagStr string
	if tagStr = tag.Get("field"); len(tagStr) != 0 {
		// ignore it if it starts with dash
		if tagStr[0] == "-"[0] {
			return nil
		}
		if p = IndexOfChar(tagStr, ","); p != -1 {
			if p > 1 {
				str := tagStr[:p]
				return &str
			}
		} else {
			return &tagStr
		}
	}

	if tagStr = tag.Get("json"); len(tagStr) == 0 {
		return nil
	}
	if tagStr[0] == "-"[0] {
		return nil
	}
	if p = IndexOfChar(tagStr, ","); p != -1 {
		if p > 1 {
			str := tagStr[:p]
			return &str
		}
		return nil
	}
	return &tagStr
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
