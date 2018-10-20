package main

import (
	_ "igo/routers"

	"igo/bootstrap"
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
