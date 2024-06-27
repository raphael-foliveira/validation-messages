package validation

import "fmt"

// fieldValueMustBe is a helper function that helps to format
// the error message for a field.
func fieldValueMustBe(s string) string {
	return fmt.Sprintf("field value must be %s", s)
}

// errValidationMessagesMap is a map of validation tags to error messages.
var errValidationMessagesMap = map[string]string{
	"required":  "field is required",
	"min":       "minimum field length is %s",
	"max":       "maximum field length is %s",
	"lt":        fieldValueMustBe("less than %s"),
	"lte":       fieldValueMustBe("less than or equal to %s"),
	"gt":        fieldValueMustBe("greater than %s"),
	"gte":       fieldValueMustBe("greater than or equal to %s"),
	"lowercase": fieldValueMustBe("all lowercase"),
	"email":     fieldValueMustBe("a valid email"),
	"password":  "password must contain at least one uppercase letter, one lowercase letter, one number, and one special character",
}

// AddValidationMessages adds a custom entry to errValidationMessagesMap.
// It also effectively enables the user to override the existing messages.
func AddValidationMessage(tag, message string) {
	errValidationMessagesMap[tag] = message
}
