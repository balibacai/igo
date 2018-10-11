package main

import (
	_ "beego/routers"

	"beego/bootstrap"
	"github.com/astaxie/beego"
)

func init() {
	bootstrap.AppConfig()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
