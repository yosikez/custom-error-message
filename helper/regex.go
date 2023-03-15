package helper

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CheckDiveField(err validator.FieldError, parent interface{}, child interface{}) (map[string]string, bool) {
	pattern := `^(\w+)\.(\w+)\[(\d+)\]\.(\w+)$`
	pattern2 := `^(\w+)\.(\w+)\.(\w+)$`
	str := err.Namespace()

	if matches, err := regexp.MatchString(pattern, str); err != nil && matches {
		re := regexp.MustCompile(pattern)

		result := re.FindStringSubmatch(str)
		mapStr := make(map[string]string)

		fieldJson, _ := reflect.TypeOf(parent).FieldByName(result[2])
		field := GetJSONTagName(fieldJson)

		attributJson, _ := reflect.TypeOf(child).FieldByName(result[4])
		attrField := GetJSONTagName(attributJson)

		mapStr["field"] = field
		mapStr["index"] = result[3]
		mapStr["attribute"] = attrField

		return mapStr, true
	} else if matches, err := regexp.MatchString(pattern2, str); err != nil && matches {
		re := regexp.MustCompile(pattern2)

		result := re.FindStringSubmatch(str)
		mapStr := make(map[string]string)

		fieldJson, _ := reflect.TypeOf(parent).FieldByName(result[2])
		field := GetJSONTagName(fieldJson)

		attributJson, _ := reflect.TypeOf(child).FieldByName(result[3])
		attrField := GetJSONTagName(attributJson)

		mapStr["field"] = field
		mapStr["attribute"] = attrField
		
		return mapStr, true
	} else {
		return nil, false
	}
}
