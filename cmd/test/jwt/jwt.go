package main

import (
	"fmt"

	"github.com/1348453525/user-redeem-code-gin/pkg/jwt"
)

func main() {
	token, _ := jwt.GenerateToken(1)
	fmt.Printf("token: %s", token)
	fmt.Println()

	claims, err := jwt.ParseToken(token + "1")
	if err != nil {
		fmt.Printf("parse token error: %v", err.Error())
	}
	fmt.Printf("claims: %+v", claims)
	fmt.Println()
}
