package main

import (
	_ "auth/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	beego.SetLogger("file", `{"filename":"./log/auth.log"}`)
	beego.SetLogFuncCall(true)
	beego.BeeLogger.DelLogger("console")

	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlDb := beego.AppConfig.String("mysqldb")
	mysqlMaxConn, err := beego.AppConfig.Int("mysqlmaxconn")
	if err != nil {
		beego.Debug("beego.AppConfig.Int(\"mysqlmaxconn\") failed, use default 30")
		mysqlMaxConn = 30
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlHost+")/"+mysqlDb+"?charset=utf8")
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", mysqlMaxConn)
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
