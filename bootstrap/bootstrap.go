package bootstrap

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"beego/extensions"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

func AppConfig()  {
	mysqlUser := beego.AppConfig.String("mysql.user")
	mysqlPass := beego.AppConfig.String("mysql.pass")
	mysqlHost := beego.AppConfig.String("mysql.host")
	mysqlDatabase := beego.AppConfig.String("mysql.database")
	ds := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlDatabase)

	orm.RegisterDataBase("default", "mysql", ds)
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = beego.AppConfig.DefaultBool("orm.debug", false)

	extensions.SetJWTSecret(beego.AppConfig.String("jwt.secret"))

	// logs config
	logs.SetLogger(beego.AppConfig.String("log.driver"))
	logs.SetLevel(beego.AppConfig.DefaultInt("log.level", logs.LevelInfo))
}