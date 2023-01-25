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
