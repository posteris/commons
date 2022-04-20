# Custom validate

Project dedicated to stadardize somo patterns like erros and validations that can be utils and shared by many others projects.

The next sessions was dedicated to show how to use each project features.

## Validation

The validation feature, actuellement have just one feature that is the ValidationModel.

### ValidateModel 

The validation model receives a model (struct) that contains the validation tag. In case of some validation fail the return will be a ValidationError array, other else nil.

```go
import commom "github.com/posteris/commons/validation"

//define struct model
type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

//create a model based on struct
userModel := User{
    Surname: "da Silva"
    Email: "dasilva@gmail.com"
}

err := commom.ValidateModel(userModel)
```

## Errors

The error package has 2 possible erros that can be understand bellow.

### DefaultError

the default error is user to return to the client a json with just one error message.

```go
import commom "github.com/posteris/commons/errors"

err := commom.CreateDefaultError("some error message")
```
### ValidationError

The validationError have 2 fields  the are showed bellow:
* __Field:__ the validated field
* __Message:__ the error message.
```go
import commom "github.com/posteris/commons/errors"

err := commom.CreateValidationError("field-name", "some error message")
```

## Query Parameter

Set of functions to operate over query parameter

### IsAsyncRequest

check if the async parameter is sent as query parameter

```go
async := IsAsyncRequest(fiberCtx)

```