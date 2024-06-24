package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	TestStringField string `validate:"required,min=3,max=10" json:"test_string_field"`
	TestIntField    int    `validate:"required,min=3,max=10" json:"test_int_field"`
}

var validStruct = &TestStruct{
	TestStringField: "valid",
	TestIntField:    5,
}

func setUp(t *testing.T) *assert.Assertions {
	assertions := assert.New(t)
	return assertions
}

func TestValidation_Success(t *testing.T) {
	tests := []struct {
		name       string
		testStruct *TestStruct
		contains   []string
	}{
		{
			name: "should return length errors when the struct has fields with invalid length",
			testStruct: &TestStruct{
				TestStringField: "a",
				TestIntField:    2,
			},
			contains: []string{
				"test_string_field",
				"test_int_field",
				"length",
				"3",
			},
		},
		{
			name:       "should return required error when the struct is missing required fields",
			testStruct: &TestStruct{},
			contains: []string{
				"test_string_field",
				"test_int_field",
				"required",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := setUp(t)
			err := Validate(tt.testStruct)
			assertions.Error(err)

			for _, c := range tt.contains {
				assertions.Contains(err.Error(), c)
			}
		})
	}
}
