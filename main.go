package main

import (
	_ "beego/routers"

	"beego/extensions"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	mysqlHost := beego.AppConfig.String("mysqlHost")
	mysqlDatabase := beego.AppConfig.String("mysqlDatabase")
	ds := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlDatabase)

	orm.RegisterDataBase("default", "mysql", ds)
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = beego.AppConfig.DefaultBool("orm.debug", false)

	extensions.SetJWTSecret(beego.AppConfig.String("jwt.secret"))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
