package models

import (
	"github.com/astaxie/beego"
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_"github.com/go-sql-driver/mysql"
)

var redisPool *redis.Pool

func init() {
	initMySql()
	initRedis()
	initUser()
	initArticle()
}

func initRedis() {
	cFile, err := goconfig.LoadConfigFile("conf/redis.conf")
	if err != nil {
		beego.Error(err)
		return
	}
	network := cFile.MustValue("", "redisnetwork", "tcp")
	address := cFile.MustValue("", "redisaddress", "127.0.0.1:6379")
	redisPool = redis.NewPool(func() (conn redis.Conn, err error) {
		conn, err = redis.Dial(network, address)
		return
	}, 30)
}

func initMySql() {
	cFile, err := goconfig.LoadConfigFile("conf/mysql.conf")
	if err != nil {
		beego.Error(err)
		return
	}

	mysqlUser := cFile.MustValue("", "mysqluser", "root")
	mysqlPwd := cFile.MustValue("", "mysqlpwd", "")
	mysqlHost := cFile.MustValue("", "mysqlhost", "localhost")
	mysqlPort := cFile.MustValue("", "mysqlport", "3306")
	mysqlDb := cFile.MustValue("", "mysqldb", "test")

	beego.Debug(mysqlUser, mysqlPwd, mysqlHost, mysqlPort, mysqlDb)

	//数据库初始化
	orm.RegisterModel(new(Article), new(Category), new(Tag), new(Reply), new(User), new(ArticleTag))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8&loc=Asia%2FChongqing", 30, 200)
	orm.RunSyncdb("default", false, true)
}