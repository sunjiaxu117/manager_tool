package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	fmt.Println(claims)
	return token.SignedString(mySigningKey)
}

func ParesToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fields := map[string]interface{}{
			"错误原因": err.Error(),
		}
		log.Println(fields)
		return nil, err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
