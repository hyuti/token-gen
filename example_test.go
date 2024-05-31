package tokengen_test

import (
	"fmt"
	"github.com/hyuti/tokengen"
	"log"
	"time"
)

func ExampleMakeToken() {
	data := "data"
	// this key should be stores in enviroment variables or something similar and only accessible by you
	secretKey := "a random key"
	token := tokengen.MakeToken(data, secretKey)
	fmt.Println(token)
}

func ExampleValidateToken() {
	data := "data"
	secretKey := "a random key"
	token := tokengen.MakeToken(data, secretKey)
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		log.Fatalln(err)
	}
	if err := tokengen.ValidateToken(data, secretKey, token, timeout); err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("token valid")
	// Output: token valid
}

func ExampleMakeTokenWithSalt() {
	data := "data"
	secretKey := "a random key"
	salt := "your own key salt"
	token := tokengen.MakeTokenWithSalt(salt, data, secretKey)
	fmt.Println(token)
}

func ExampleValidateTokenWithKeySalt() {
	data := "data"
	secretKey := "a random key"
	salt := "your own key salt"
	token := tokengen.MakeTokenWithSalt(salt, data, secretKey)
	timeout, err := time.ParseDuration("60s")
	if err != nil {
		log.Fatalln(err)
	}
	if err := tokengen.ValidateTokenWithKeySalt(salt, data, secretKey, token, timeout); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("token valid")
	// Output: token valid
}
