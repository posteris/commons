# Custom validate

This code was created to standardize the API error output for all microservices



```go
import commom "github.com/posteris/commom-validate"

type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

userModel := User{
    Surname: "da Silva"
    Email: "dasilva@gmail.com"
}

err := commom.ValidateModel(userModel)
```


```go
import commom "github.com/posteris/commom-validate"

err := commom.CreateDefaultError(userModel)
```