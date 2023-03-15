package custom_error_message

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/yosikez/custom-error-message/helper"
)

func GetErrMess(err error, parentStruct interface{}, childStruct interface{}) map[string]string {
	errFields := make(map[string]string)
	var errs validator.ValidationErrors

	if errors.As(err, &errs) {
		for _, errField := range errs {
			var field, disField string

			mapStr, isTrue := helper.CheckDiveField(errField, parentStruct, childStruct)

			if isTrue {
				_, ok := mapStr["index"]
				if ok {
					field = fmt.Sprintf("%s.%s.%s", mapStr["field"], mapStr["index"], mapStr["attribute"])
				} else {
					field = fmt.Sprintf("%s.%s", mapStr["field"], mapStr["attribute"])
				}
				disField = mapStr["attribute"]
			} else {
				structField, _ := reflect.TypeOf(parentStruct).FieldByName(errField.Field())
				field = helper.GetJSONTagName(structField)
				disField = field
			}

			switch errField.Tag() {
			case "required":
				errFields[field] = fmt.Sprintf("%s is required", disField)
			case "email":
				errFields[field] = fmt.Sprintf("%s is not a valid email address", disField)
			case "min":
				errFields[field] = fmt.Sprintf("%s must be at least %s characters/nums long", disField, errField.Param())
			case "max":
				errFields[field] = fmt.Sprintf("%s must be at most %s characters/nums long", disField, errField.Param())
			case "gender":
				errFields[field] = fmt.Sprintf("%s must be male or female", disField)
			case "uniqueMail":
				errFields[field] = fmt.Sprintf("%s already taken", disField)
			case "numeric":
				errFields[field] = fmt.Sprintf("%s must be numeric", disField)
			case "datetime":
				errFields[field] = fmt.Sprintf("%s format must be yyyy-mm-dd hh:mm:ss", disField)
			case "uniqueField":
				errFields[field] = fmt.Sprintf("%s already taken", disField)
			default:
				errFields[field] = fmt.Sprintf("%s is not valid", disField)
			}
		}
	} else {
		var unmarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalErr) {
			errFields[unmarshalErr.Field] = fmt.Sprintf("Invalid type. Expected %v but got %v", unmarshalErr.Type, unmarshalErr.Value)
		} else {
			errFields["input"] = err.Error()
		}
	}

	return errFields
}