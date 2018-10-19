package controllers

import (
	"github.com/astaxie/beego"
	"beego/models"
	"time"
	"beego/response"
	"beego/extensions"
	"beego/filters"
	"strings"
	"github.com/astaxie/beego/logs"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
	user    *models.User
	isLogin bool
}


func (this *baseController) Prepare() {
	// page start time
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	//this.Data["AppDescription"] = utils.AppDescription

	this.initAuth()

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

// init current user with token from http request headers
func (this *baseController) initAuth() {
	this.isLogin = false

	// get token from Authorization Header
	tokenFullString := this.Ctx.Input.Header("Authorization")

	if !strings.HasPrefix(tokenFullString, "Bearer ") {
		this.JsonOutput(response.JsonResult{Error: 101001, Msg: "require token"})
	}

	tokenString := tokenFullString[7:]

	//fmt.Println("token:" + tokenString)

	// parse token with claims
	token, err := extensions.ParseRSAJWTTokenWithClaims(tokenString, &filters.LoginClaims{})

	if err != nil {
		logs.Error(err)
		this.JsonOutput(response.JsonResult{Error: 101002, Msg: "parse token error"})
	}

	// validate & extract token
	if claims, ok := token.Claims.(*filters.LoginClaims); ok && token.Valid {
		user, err := models.GetUserById(claims.UserID)
		if err != nil {
			this.JsonOutput(response.JsonResult{Error: 101004, Msg: "user not exist"})
		}
		// assign user
		this.user = user
		this.isLogin = true
	} else {
		this.JsonOutput(response.JsonResult{Error: 101003, Msg: "token expired"})
	}
}

// output json data and exit
func (this *baseController) JsonOutput(data interface{})  {
	this.Data["json"] = &data
	this.ServeJSON()
	this.StopRun()
}

