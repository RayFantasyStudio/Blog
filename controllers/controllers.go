package controllers

import "github.com/astaxie/beego"

var owner string

func init() {
	owner = beego.AppConfig.String("owner")
}

func initHeaderFooterData(c *beego.Controller, title string) {
	c.Data["Owner"] = owner

	// TODO: 实现登陆状态检查
	c.Data["User"] = "访客"

	c.Data["Title"] = title
}