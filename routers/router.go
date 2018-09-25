// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func init() {

	mySigningKey := []byte(beego.AppConfig.String("jwt.secret"))

	type LoginClaims struct {
		UserID int64
		jwt.StandardClaims
	}

	// filters
	var auth = func(ctx *context.Context) {
		tokenString := ctx.Input.Header("token")
		//fmt.Println("token:" + tokenString)

		if len(tokenString) == 0 && ctx.Request.RequestURI != "/v1/login" {
			ctx.Redirect(302, "/login")
		}

		// parse token with claims
		token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Println(err)
			ctx.Redirect(302, "/login")

		}

		// validate & extract token
		if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Println(err)
			ctx.Redirect(302, "/login")
		}
	}

	// routers
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(auth),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
