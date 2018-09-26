package filters

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/astaxie/beego/context"
	"beego/extensions"
)

var (
	authExcept = map[string]bool{"/v1/login": true}
)

type LoginClaims struct {
	UserID int64
	jwt.StandardClaims
}


func Auth(ctx *context.Context) {
	tokenString := ctx.Input.Header("token")
	//fmt.Println("token:" + tokenString)

	_, isExcept := authExcept[ctx.Request.RequestURI]
	if isExcept {
		return
	}

	//if len(tokenString) == 0 {
	//	ctx.Redirect(302, "/login")
	//}

	// parse token with claims
	token, err := extensions.ParseJWTTokenWithClaims(&LoginClaims{}, tokenString)

	if err != nil {
		fmt.Println(err)
		ctx.Redirect(302, "/login")
		return
	}

	// validate & extract token
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
		ctx.Redirect(302, "/login")
	}
}