package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"beego/extensions"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func init() {
	extensions.SetJWTSecret("123")
	extensions.SetPrivateKey("../rsaprivatekey.pem")
	extensions.SetPublicKey("../rsapublickey.pem")
}

type LoginClaims struct {
	UserID int64
	jwt.StandardClaims
}

// TestGet is a sample to run an endpoint test
func TestJWTBuildAndParseToken(t *testing.T) {
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
	now := time.Now()
	userID := int64(1234567)
	expiredAt := now.Unix() + 3600

	tokenString, err := extensions.NewRSAJWTTokenStringWithClaims(LoginClaims{
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
	token, err := extensions.ParseRSAJWTTokenWithClaims(tokenString, &LoginClaims{})

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

