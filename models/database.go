package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)
const (
)
func init(){
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPwd := beego.AppConfig.String("mysqlpwd")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlPort := beego.AppConfig.String("mysqlport")
	mysqlDb := beego.AppConfig.String("mysqldb")

	//数据库初始化
	orm.RegisterModel(new(Article))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",
		mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8&loc=Asia%2FChongqing", 30, 200)
	orm.RunSyncdb("default", false, true)

}