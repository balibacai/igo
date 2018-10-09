package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/logs"
	"time"
	"beego/extensions"
	"github.com/dgrijalva/jwt-go"
	"beego/filters"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
}

type LoginCredentials struct {
	Email string		`form:"email"valid:"Required;Email;MaxSize(32)"`
	Password string		`form:"password"valid:"Required;MaxSize(64)"`
}

type JsonResult struct {
	Error int32			`json:"error"`
	Msg string 			`json:"msg"`
	Data interface{}	`json:"data"`
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {

	var result JsonResult

	// login credentials
	credentials := LoginCredentials{}
	if err := c.ParseForm(&credentials); err != nil {
		result = JsonResult{Error: 100001, Msg: "error occurs when parsing login form"}
		logs.Error(&result, err)
		c.Data["json"] = &result
		c.ServeJSON()
		return
	}

	logs.Debug("login from data", credentials)

	// validate
	valid := validation.Validation{}
	passed, err := valid.Valid(&credentials)
	if err != nil {
		result = JsonResult{Error: 100002, Msg: "error occurs when validating login credentials"}
		logs.Error(&result, err)
		c.Data["json"] = &result
		c.ServeJSON()
		return
	}

	if !passed {
		result = JsonResult{Error: 100003, Msg: "login credentials invalid"}
		logs.Error(&result, valid.ErrorsMap)
		c.Data["json"] = &result
		c.ServeJSON()
		return
	}

	// attempt login
	now := time.Now()
	userID := int64(1234567)
	// expired after 30 days
	expiredAt := now.Unix() + 2592000

	tokenString, err := extensions.NewJWTTokenStringWithClaims(filters.LoginClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expiredAt,
			Issuer: "igo",
		},
	})

	if err != nil {
		result = JsonResult{Error: 100004, Msg: "error occurs when generating login token"}
		logs.Error(&result, err)
		c.Data["json"] = &result
		c.ServeJSON()
		return
	}

	tokenMap := map[string]interface{}{
		"token": tokenString,
		"expiredAt": expiredAt,
	}

	// return token
	result = JsonResult{Error: 0, Data: &tokenMap}
	c.Data["json"] = &result
	c.ServeJSON()
}