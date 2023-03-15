# Custom Error Message

The Custom Error Message is a Go module that provides a utility function for creating custom error messages based on the given error, parent struct, and child struct. This module is very useful for validating error structs in Go.

## Installation

To use this package, you need to have Go installed on your system.

To install this module, run the following command in your terminal:

```bash
go get github.com/yosikez/custom-error-message
```

Then, import the package into your Go code:
```go
import "github.com/yosikez/custom-error-message"
```

## Usage

The GetErrMess function is used to create custom error messages. This function accepts three parameters, namely the error to be modified, the parent struct, and the child struct. The parent struct is required to validate the error struct.

The child struct can be ignored by passing a nil value to that parameter.

Here's an example usage of the GetErrMess function in gin framework:


### Example 1
```golang
package controller

import (
    "fmt"

    cusMessage "github.com/yosikez/custom-error-message"
)

type Person struct {
	Name    string    `json:"name" binding:"required"`
	Address []Address `json:"address" binding:"required,dive"`
}

type Address struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
}


func Create(c *gin.Context) {
    var person Person

	if err := c.ShouldBind(&person); err != nil {
		errFields := cusMessage.GetErrMess(err, person, Address{})

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"errors":  errFields,
		})
		return
	}

    // your code
}
```
### Output example 1
```json
{
    "message" : "validation error",
    "errors" : {
        "name" : "name is required.",
        "address.0.street" : "street is required.",
        "address.0.city": "city is required."
    }
}

```


### Example 2

```golang
package controller

import (
    "fmt"

    cusMessage "github.com/yosikez/custom-error-message"
)

type Person struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

func Create(c *gin.Context) {
    var person Person

	if err := c.ShouldBind(&person); err != nil {
		errFields := cusMessage.GetErrMess(err, person, nil)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"errors":  errFields,
		})
		return
	}

    // your code
}

```
### Output example 2
```json
{
    "message" : "validation error",
    "errors" : {
        "name" : "name is required.",
    	"email" : "email is required.",
    	"password": "password is required."
    }
}

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

