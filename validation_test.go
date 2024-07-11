package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) *assert.Assertions {
	assertions := assert.New(t)
	return assertions
}

func TestValidation_InvalidStruct(t *testing.T) {
	type TestStruct struct {
		TestStringField string `validate:"required,min=3,max=10" json:"test_string_field"`
		TestIntField    int    `validate:"required,min=3,max=10" json:"test_int_field"`
	}

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

func TestValidation_ValidStruct(t *testing.T) {
	type TestStruct struct {
		TestStringField string `validate:"required,min=3,max=10" json:"test_string_field"`
		TestIntField    int    `validate:"required,min=3,max=10" json:"test_int_field"`
	}

	tests := []struct {
		testStruct *TestStruct
		name       string
	}{
		{
			name: "should return nil when the struct has valid fields",
			testStruct: &TestStruct{
				TestStringField: "abc",
				TestIntField:    3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertions := setUp(t)
			err := Validate(tt.testStruct)
			assertions.Nil(err)
		})
	}
}
