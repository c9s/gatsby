package sqlutils

import (
	"reflect"
	"strings"
)

// Extract column name attribute from struct tag (the first element) of the 'field' tag or
// column name from 'json' tag.
func GetColumnNameFromTag(tag *reflect.StructTag) *string {

	if tagStr := tag.Get("field"); tagStr != "" {
		// ignore it if it starts with dash
		if strings.HasPrefix(tagStr, "-") {
			return nil
		}
		fieldTags := strings.SplitN(tagStr, ",", 2)
		if fieldTags[0] != "" {
			return &fieldTags[0]
		}
	}

	if jsonTagStr := tag.Get("json"); jsonTagStr != "" {
		if strings.HasPrefix(jsonTagStr, "-") {
			return nil
		}
		jsonTags := strings.SplitN(jsonTagStr, ",", 2)
		if jsonTags[0] != "" {
			return &jsonTags[0]
		}
	}
	return nil
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
