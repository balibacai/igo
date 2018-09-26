package extensions

import (
	"github.com/dgrijalva/jwt-go"
)

var (
	myJwt = new(JWT)
)

type JWT struct {
	Secret string
}

func (j *JWT) SetSecret(secret string)  {
	j.Secret = secret
}

func (j *JWT) ParseWithClaims(claims jwt.Claims, tokenString string) (*jwt.Token, error) {
	// parse token with claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	return token, err
}

func SetJWTSecret(secret string)  {
	myJwt.Secret = secret
}

func ParseJWTTokenWithClaims(claims jwt.Claims, tokenString string) (*jwt.Token, error) {
	// parse token with claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return myJwt.Secret, nil
	})

	return token, err
}
