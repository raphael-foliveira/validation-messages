# Easy JSON formatted validation messages

## Installation

```bash
go get github.com/raphael-foliveira/validation
```

## Usage

```go
type User struct {
 // The validate tag is used to define the
 // validation rules for the field.
 // If a json tag is provided, the field name
 //in the ValidationError will be the json tag.
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
 fmt.Println(err) // {"email":["field value must be a valid email"],"password":["minimum field length is 8"],"username":["minimum field length is 6"]}
}
```

PS: Only structs can be validated.
