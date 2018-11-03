package web

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()
var jsonTagRegex = regexp.MustCompile(`^([^,]*)*`)

const gtinFormat = `^\d{14}$`

func init() {
	// make it so the validation prefers using the `json` tag to fetch the field name
	validate.RegisterTagNameFunc(getValidationFieldName)

}

// Invalid describes a validation error belonging to a specific field.
type Invalid struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InvalidError is a custom error type for invalid fields. Backwards compatible.
type InvalidError []Invalid

// Error implements the error interface for InvalidError.
func (err InvalidError) Error() string {
	var str string
	for _, v := range err {
		str = fmt.Sprintf("%s,{%s:%s}", str, v.Field, v.Error)
	}
	return strings.TrimLeft(str, ",")
}

func getValidationFieldName(field reflect.StructField) string {
	value := field.Tag.Get("json")
	// if the JSON tag is to be ignored, return ""
	if value == "-" {
		value = ""
	}
	// if we actually have a JSON tag to work with, only grab the value before the ,
	if value != "" {
		values := jsonTagRegex.FindStringSubmatch(value)
		value = values[1]
	}
	if value == "" {
		value = field.Name
	}
	// return
	return value
}

// Validate wraps the validation logic to fit our needs.
func Validate(v interface{}) error {
	var inv InvalidError
	if fve := validate.Struct(v); fve != nil {
		for _, fe := range fve.(validator.ValidationErrors) {
			inv = append(inv, Invalid{Field: fe.Field(), Error: fe.Tag()})
		}
		return inv
	}

	return nil
}
