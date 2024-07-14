package validation

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// getValidationMessage returns the validation message for a field error.
// If the tag is not found in the map, it returns the default error message.
// If the tag is found in the map, it returns the message with the param if it exists.
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

// ErrorMap is a map of field names to a list of validation error messages.
// The field name is the json tag if it exists, otherwise it is the struct field's name.
type ErrorMap map[string][]string

// Error returns a string representation of the ValidationError.
// It marshals the map to json and returns the string.
func (ve ErrorMap) Error() string {
	errJson, err := json.Marshal(ve)
	if err != nil {
		return fmt.Sprintf("error marshalling validation error: %s", err.Error())
	}
	return string(errJson)
}

func buildErrorMap(strct interface{}, validationError validator.ValidationErrors) ErrorMap {
	errMap := ErrorMap{}

	strctType := reflect.TypeOf(strct).Elem()

	for _, fieldError := range validationError {
		fieldName := fieldError.Field()

		field, found := strctType.FieldByName(fieldName)
		if !found {
			fmt.Printf("field %s not found in struct\n", fieldError.Field())
			continue
		}

		jsonName, jsonSet := field.Tag.Lookup("json")

		if jsonSet {
			fieldName = jsonName
		}

		validationErrorMessage := getValidationMessage(fieldError)
		errMap[fieldName] = append(errMap[fieldName], validationErrorMessage)
	}

	return errMap
}

// Validate validates a struct using the go-playground/validator package.
// It returns a ValidationError if there are any validation ValidationErrors.
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

	return buildErrorMap(strct, validationError)
}
