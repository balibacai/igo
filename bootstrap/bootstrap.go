package bootstrap

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mistcheng/ilib/ijwt"
)

func AppConfig() {
	// config jwt
	ijwt.SetJWTMode(ijwt.JWTMode(beego.AppConfig.DefaultInt("jwt.mode", 0)))
	if jwtSecret := beego.AppConfig.String("jwt.secret"); len(jwtSecret) > 0 {
		ijwt.SetJWTSecret(jwtSecret)
	}

	if jwtPublicKeyPath := beego.AppConfig.String("jwt.public_key_pem_path"); len(jwtPublicKeyPath) > 0 {
		ijwt.SetJWTPublicKey(jwtPublicKeyPath)
	}

	// logs config
	logs.SetLogger(beego.AppConfig.String("log.driver"))
	logs.SetLevel(beego.AppConfig.DefaultInt("log.level", logs.LevelInfo))
}
