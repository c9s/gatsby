package sqlutils

import "reflect"
import "strings"

// Extract column name attribute from struct tag (the first element) of the 'field' tag or
// column name from 'json' tag.
func GetColumnNameFromTag(tag *reflect.StructTag) *string {
	fieldTags := strings.Split(tag.Get("field"), ",")
	if len(fieldTags[0]) > 0 {
		return &fieldTags[0]
	}
	jsonTags := strings.Split(tag.Get("json"), ",")
	if len(jsonTags[0]) > 0 {
		return &jsonTags[0]
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
