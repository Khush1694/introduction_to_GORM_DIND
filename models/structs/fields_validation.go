package structs

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// ValidateFields validates all the fields of a model
func ValidateFields(m interface{}) Response {
	validations := []Validation{}

	if err := validator.New().Struct(m); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, v := range errs {
			validations = append(validations, Validation{
				Field:   v.Field(),
				Message: generateValidationMessage(v.Field(), v.Tag()),
			})
		}

		return Response{
			Success:     false,
			Validations: validations,
		}
	}

	return Response{}
}

// GenerateValidationMessage generates a validation message for each wrong field
func generateValidationMessage(field, rule string) string {
	switch rule {
	case "required":
		return fmt.Sprintf("Field %s is %s ", field, rule)
	default:
		return fmt.Sprintf("Field %s is no valid", field)
	}
}
