package examples

import (
	"fmt"

	"github.com/raphael-foliveira/validation-messages"
)

type User struct {
	Username string `validate:"required,min=6,max=32" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=32" json:"password"`
}

func ValidateStruct() {
	myStruct := &User{
		Username: "inv",
		Email:    "inv.com",
		Password: "inv",
	}

	err := validation.Validate(myStruct)
	fmt.Println(err)
}
