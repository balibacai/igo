package filters

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/astaxie/beego/context"
	"beego/extensions"
	"beego/response"
)

var (
	authExcept = map[string]bool{"/v1/login": true}
)

type LoginClaims struct {
	UserID int64
	jwt.StandardClaims
}


func Auth(ctx *context.Context) {

	var result response.JsonResult

	tokenString := ctx.Input.Header("token")
	//fmt.Println("token:" + tokenString)

	_, isExcept := authExcept[ctx.Request.URL.Path]
	if isExcept {
		return
	}

	if len(tokenString) == 0 {
		result = response.JsonResult{Error: 101001, Msg: "require token"}
		ctx.Output.JSON(&result, false, false)
		return
	}

	// parse token with claims
	token, err := extensions.ParseJWTTokenWithClaims(tokenString, &LoginClaims{})

	if err != nil {
		result = response.JsonResult{Error: 101002, Msg: "parse token error"}
		ctx.Output.JSON(&result, false, false)
		return
	}

	// validate & extract token
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
	} else {
		result = response.JsonResult{Error: 101003, Msg: "token expired"}
		ctx.Output.JSON(&result, false, false)
		return
	}
}