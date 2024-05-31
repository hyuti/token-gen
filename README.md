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
# Examples
See [examples](https://github.com/hyuti/token-gen/blob/main/example_test.go) for further details

# Author
**Hyuti Le**
* https://github.com/hyuti
