package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func getJSONFieldName(structType any, fieldName string) string {
	t := reflect.TypeOf(structType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if field, ok := t.FieldByName(fieldName); ok {
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			return jsonTag
		}
	}
	return fieldName
}

func AppValidator(req any, c *fiber.Ctx) *[]ErrorDetail {
	if err := validate.Struct(req); err != nil {
		var errors []ErrorDetail

		for _, err := range err.(validator.ValidationErrors) {
			msg := "Invalid or missing value"
			if err.Tag() == "required" {
				msg = "This field is required"
			} else if err.Tag() == "email" {
				msg = "Invalid email format"
			} else if err.Tag() == "gt" {
				msg = "Cannot be empty"
			}
			fieldName := getJSONFieldName(req, err.StructField())
			errors = append(errors,
				ErrorDetail{
					Field: fieldName,
					Issue: msg,
				})
		}

		return &errors
	}

	return nil
}
