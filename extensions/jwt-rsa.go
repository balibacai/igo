package extensions

import (
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"io/ioutil"
)

var (
	myRSAJwt = new(RSAJWT)
)

type RSAJWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func (j *RSAJWT) SetPrivateKey(pem []byte)  {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		panic(err)
	}
	j.PrivateKey = privateKey
}

func (j *RSAJWT) SetPublicKey(pem []byte)  {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	if err != nil {
		panic(err)
	}
	j.PublicKey = publicKey
}

func (j *RSAJWT) ParseWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	// parse token with claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})

	return token, err
}

func (j *RSAJWT) NewSignatureWithClaims(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(j.PrivateKey)
	return ss, err
}

func SetPublicKey(path string)  {
	pemBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	myRSAJwt.SetPublicKey(pemBytes)
}

func SetPrivateKey(path string)  {
	pemBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	myRSAJwt.SetPrivateKey(pemBytes)
}

func ParseRSAJWTTokenWithClaims(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return myRSAJwt.ParseWithClaims(tokenString, claims)
}

func NewRSAJWTTokenStringWithClaims(claims jwt.Claims) (string, error) {
	return myRSAJwt.NewSignatureWithClaims(claims)
}

