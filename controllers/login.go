package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"beego/response"
	"github.com/astaxie/beego/httplib"
	"fmt"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (this *LoginController) Post() {

	var result response.JsonResult

	req := httplib.Post(fmt.Sprintf("%s/v1/login", beego.AppConfig.String("auth.url")))
	req.Param("email", this.GetString("email", ""))
	req.Param("password", this.GetString("password", ""))
	req.SetTimeout(time.Second / 2, time.Second)
	req.ToJSON(&result)

	this.Data["json"] = &result
	this.ServeJSON()
}