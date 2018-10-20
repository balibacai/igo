package main

import (
	_ "igo/routers"

	"github.com/astaxie/beego"
	"igo/bootstrap"
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
