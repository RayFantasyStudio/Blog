package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)
const (
	_DB_USER_NAME = beego.AppConfig.String("mysqluser")
	_DB_PWD = beego.AppConfig.String("mysqlpwd")
	_DB_URL = beego.AppConfig.String("mysqlurl")
	_DB_PORT = beego.AppConfig.String("mysqlport")
	_DB_NAME = beego.AppConfig.String("mysqldb")
)
func init(){
	//数据库初始化
	orm.RegisterModel(new(Article))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",
		_DB_USER_NAME + ":" + _DB_PWD + "@tcp(" + _DB_URL + ":" + _DB_PORT + ")/" + _DB_NAME + "?charset=utf8&loc=Asia%2FChongqing", 30, 200)

}