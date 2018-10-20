package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"igo/extensions"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func init() {
	extensions.SetJWTSecret("123")
	extensions.SetJWTPrivateKey("../rsaprivatekey.pem")
	extensions.SetJWTPublicKey("../rsapublickey.pem")
}

type LoginClaims struct {
	UserID int64
	jwt.StandardClaims
}

// TestGet is a sample to run an endpoint test
func TestJWTBuildAndParseToken(t *testing.T) {
	extensions.SetJWTMode(extensions.JWTSecretMode)

	now := time.Now()
	userID := int64(1234567)
	expiredAt := now.Unix() + 3600

	tokenString, err := extensions.NewJWTTokenStringWithClaims(LoginClaims{
		userID,
		jwt.StandardClaims {
			ExpiresAt: expiredAt,
			Issuer: "test",
		},
	})

	Convey("build err should be nil", t, func() {
		So(err, ShouldBeNil)
	})


	// parse token with claims
	token, err := extensions.ParseJWTTokenWithClaims(tokenString, &LoginClaims{})

	Convey("parse err should be nil", t, func() {
		So(err, ShouldBeNil)
	})

	claims, ok := token.Claims.(*LoginClaims)

	Convey("token check", t, func() {
		So(ok, ShouldBeTrue)
		So(token.Valid, ShouldBeTrue)
	})

	Convey("claims check", t, func() {
		So(claims.UserID, ShouldEqual, userID)
		So(claims.StandardClaims.ExpiresAt, ShouldEqual, expiredAt)
	})
}

func TestRSAJWTBuildAndParseToken(t *testing.T) {
	extensions.SetJWTMode(extensions.JWTRSAMode)

	now := time.Now()
	userID := int64(1234567)
	expiredAt := now.Unix() + 3600

	tokenString, err := extensions.NewJWTTokenStringWithClaims(LoginClaims{
		userID,
		jwt.StandardClaims {
			ExpiresAt: expiredAt,
			Issuer: "test",
		},
	})

	Convey("build err should be nil", t, func() {
		So(err, ShouldBeNil)
	})

	//fmt.Printf(tokenString)

	// parse token with claims
	token, err := extensions.ParseJWTTokenWithClaims(tokenString, &LoginClaims{})

	Convey("parse err should be nil", t, func() {
		So(err, ShouldBeNil)
	})

	claims, ok := token.Claims.(*LoginClaims)

	Convey("token check", t, func() {
		So(ok, ShouldBeTrue)
		So(token.Valid, ShouldBeTrue)
	})

	//fmt.Printf("%d", claims.UserID)

	Convey("claims check", t, func() {
		So(claims.UserID, ShouldEqual, userID)
		So(claims.StandardClaims.ExpiresAt, ShouldEqual, expiredAt)
	})
}

