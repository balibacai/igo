package extensions

import (
	"github.com/dgrijalva/jwt-go"
)

var (
	myJwt = new(JWT)
)

type JWT struct {
	Secret []byte
}

func (j *JWT) SetSecret(secret string)  {
	j.Secret = []byte(secret)
}

func (j *JWT) ParseWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	// parse token with claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	return token, err
}

func (j *JWT) NewSignatureWithClaims(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(j.Secret)

	return ss, err
}

func SetJWTSecret(secret string)  {
	myJwt.SetSecret(secret)
}

func ParseJWTTokenWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return myJwt.ParseWithClaims(tokenString, claims)
}

func NewJWTTokenStringWithClaims(claims jwt.Claims) (string, error) {
	return myJwt.NewSignatureWithClaims(claims)
}

