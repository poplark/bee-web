package main

import (
	"fmt"

	"bee-web/common"
	_ "bee-web/routers"

	"github.com/astaxie/beego"
)

func init() {

	dbUser := beego.AppConfig.String("dbUser")
	dbPass := beego.AppConfig.String("dbPass")
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbName := beego.AppConfig.String("dbName")

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("[dbDSN]: ", dbDSN)
	common.InitOrm("default", dbDSN, true)

	common.InitLog()
}

func main() {
	common.RunOrm("default")

	beego.Run("localhost")
}
