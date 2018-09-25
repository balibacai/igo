package main

import (
	_ "beego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqldb := beego.AppConfig.String("mysqldb")
	ds := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqluser, mysqlpass, mysqlhost, mysqldb)

	orm.RegisterDataBase("default", "mysql", ds)
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = beego.AppConfig.DefaultBool("orm.debug", false)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

