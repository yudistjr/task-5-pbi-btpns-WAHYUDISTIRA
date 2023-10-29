package helper

import "github.com/golang-jwt/jwt/v4"

func SetToken() (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte(RandomString(16)))
	return token, err
}