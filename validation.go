package validation

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// getValidationMessage returns the validation message for a field error.
//
// If the tag is not found in the map, it returns the default error message.
//
// If the tag is found in the map, it returns the message with the param if it exists.
//
// If the param does not exist, it returns the message without the param.
func getValidationMessage(fieldError validator.FieldError) string {
	tag := fieldError.Tag()

	message, exists := errValidationMessagesMap[tag]
	if !exists {
		return fieldError.Error()
	}

	param := fieldError.Param()
	if param != "" {
		return fmt.Sprintf(message, fieldError.Param())
	}

	return message
}

// ValidationError is a map of field names to a list of validation error messages.
//
// The field name is the json tag if it exists, otherwise it is the struct field's name.
type ValidationError map[string][]string

// Error returns a string representation of the ValidationError.
//
// It marshals the map to json and returns the string.
func (ve ValidationError) Error() string {
	errJson, err := json.Marshal(ve)
	if err != nil {
		return fmt.Sprintf("error marshalling validation error: %s", err.Error())
	}
	return string(errJson)
}

// Validate validates a struct using the go-playground/validator package.
//
// It returns a ValidationError if there are any validation ValidationErrors.
//
// If there are no ValidationErrors, it returns nil.
func Validate(strct interface{}) error {
	validate := validator.New()
	err := validate.Struct(strct)
	if err == nil {
		return nil
	}
	validationError, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	strctType := reflect.TypeOf(strct).Elem()

	validationErrorMap := ValidationError{}

	for _, fieldError := range validationError {
		field, found := strctType.FieldByName(fieldError.Field())
		if !found {
			fmt.Printf("field %s not found in struct\n", fieldError.Field())
			continue
		}

		fieldName := fieldError.Field()

		jsonName, fieldSet := field.Tag.Lookup("json")

		if fieldSet {
			fieldName = jsonName
		}

		validationErrorMessage := getValidationMessage(fieldError)
		validationErrorMap[fieldName] = append(validationErrorMap[fieldName], validationErrorMessage)
	}

	return validationErrorMap
}
