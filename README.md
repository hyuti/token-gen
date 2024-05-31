# Token Generator
TokenGen is intended to assure data integrity and authenticity shared between parties. This package is inspired by [django-contrib-auth-token](https://github.com/django/django/blob/main/django/contrib/auth/tokens.py) so that each token generated will have its timeout.
Used techniques are symmetric and hashing so that you must have a secret key shared among partners to succeed on using this module.
# Use cases
- Digital signatures
- OTP (one time password)
- Forget password token with timeout
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

	"github.com/hyuti/tokengen"
)

func main() {
	data := "data"
	// this key should be stores in enviroment variables or something similar and only accessible by you
	secretKey := "a random key"
	token := tokengen.MakeToken(data, secretKey)
	fmt.Println(token)
}

```
## Validate a token 
```go
package main

import (
	"fmt"
	"time"

	"github.com/hyuti/tokengen"
)

func main() {
	data := "data"
	secretKey := "a random key"
	token := tokengen.MakeToken(data, secretKey)
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := tokengen.ValidateToken(data, secretKey, token, timeout);err != nil {
		fmt.Printf("invalid token: %s\n", err)
		return
	}
	fmt.Println("token valid")
}

```
## Make a token with your own key salt
```go
package main

import (
	"fmt"

	"github.com/hyuti/tokengen"
)

func main() {
	data := "data"
	secretKey := "a random key"
	salt := "your own key salt"
	token := tokengen.MakeTokenWithSalt(salt, data, secretKey)
	fmt.Println(token)
}
```
## Validate a token with your own key salt
```go
package main

import (
	"fmt"
	"time"

	"github.com/hyuti/tokengen"
)

func main() {
	data := "data"
	secretKey := "a random key"
	salt := "your own key salt"
	token := tokengen.MakeTokenWithSalt(salt, data, secretKey)
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := tokengen.ValidateTokenWithKeySalt(salt, data, secretKey, token, timeout);err != nil {
		fmt.Printf("invalid token: %s\n", err)
		return
	}
	fmt.Println("token valid")
}

```
# Author
**Hyuti Le**
* https://github.com/hyuti
