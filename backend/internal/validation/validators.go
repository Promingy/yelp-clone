package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func Validate(s interface{}) map[string]string {
	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	errs := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := toSnakeCase(err.Field())
		errs[field] = formatError(err)
	}

	return errs
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func formatError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s characters", err.Param())
	case "max":
		return fmt.Sprintf("Must be no more than %s characters", err.Param())
	case "url":
		return "Must be a valid URL"
	case "omitempty":
		return "Invalid value"
	case "e164":
		return "Must be a valid phone number (e.g., +15551234567)"
	case "iso3166_1_alpha2":
		return "Must be a valid 2-letter country code (e.g., US)"
	case "len":
		return fmt.Sprintf("Must be exactly %s characters", err.Param())
	default:
		return "Invalid value"
	}
}