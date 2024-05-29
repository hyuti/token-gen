# Token gen
Token generator is a token generator used for use cases such as reset password, otp, etc. Basically, you will need this package in use cases where data shared between parties must be ensure its integrity and authenticity. This package is inspired by [django-contrib-auth-token](https://github.com/django/django/blob/main/django/contrib/auth/tokens.py) so that each token generated will have its timeout.
# Installation
Use go get
```sh 
go get -u github.com/hyuti/tokengen
``` 
Then import the package into your own code
```go
import "github.com/hyuti/tokengen"
```
# Usage
## Make a token 
```go
package main

import (
	"fmt"

	"github.com/hyuti/pwdTokenGenerator/generator"
)

func main() {
	data := "data"
	// this key should be stores in enviroment variables or something similar and only accessible by you
	secretKey := "a random key"
	token := generator.MakeToken(data, secretKey)
	fmt.Println(token)
}

```
## Validate a token 
```go
package main

import (
	"fmt"
	"time"

	"github.com/hyuti/pwdTokenGenerator/generator"
)

func main() {
	data := "data"
	secretKey := "a random key"
	token := generator.MakeToken(data, secretKey)
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Println(err)
	} else {
		err = generator.ValidateToken(data, secretKey, token, timeout)
		if err != nil {
			fmt.Printf("invalid token: %s\n", err)
		} else {
			fmt.Println("token valid")
		}
	}
}

```
## Make a token with your own key salt
```go
package main

import (
	"fmt"

	"github.com/hyuti/pwdTokenGenerator/generator"
)

func main() {
	data := "data"
	secretKey := "a random key"
	salt := "your own key salt"
	token := generator.MakeTokenWithSalt(salt, data, secretKey)
	fmt.Println(token)
}
```
## Validate a token with your own key salt
```go
package main

import (
	"fmt"
	"time"

	"github.com/hyuti/pwdTokenGenerator/generator"
)

func main() {
	data := "data"
	secretKey := "a random key"
	salt := "your own key salt"
	token := generator.MakeTokenWithSalt(salt, data, secretKey)
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Println(err)
	} else {
		err = generator.ValidateTokenWithKeySalt(salt, data, secretKey, token, timeout)
		if err != nil {
			fmt.Printf("invalid token: %s\n", err)
		} else {
			fmt.Println("token valid")
		}
	}
}

```
# Author
**Hyuti Le**
* https://github.com/hyuti
