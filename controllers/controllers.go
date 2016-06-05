package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

var owner string

func init() {
	owner = beego.AppConfig.String("owner")
}

func initHeaderFooterData(c *beego.Controller, title string) (user *models.User, err error) {
	c.Data["Owner"] = owner

	c.Data["Title"] = title

	if token := c.Ctx.GetCookie("token"); len(token) != 0 {
		if user, err = models.GetUserByToken(token); err != nil {
			beego.Error(err)
		} else {
			c.Data["IsLogin"] = user != nil
			c.Data["User"] = user
		}
	}
	return
}