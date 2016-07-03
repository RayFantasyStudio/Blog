package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"fmt"
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
			c.Data["IsLogin"] = true
			c.Data["User"] = user
			var avatar_path string
			avatar_path, err = models.GetUserAvatar(int64(user.Id))
			c.Data["avatar_path"] = avatar_path
			beego.Info(avatar_path)
		}
	}

	flash := beego.ReadFromRequest(c)
	if n, ok := flash.Data["notice"]; ok {
		c.Data["Message"] = n
	}
	return
}

func NewNoticeFlash(c *beego.Controller, msg string, args ...interface{}) {
	var data string
	if len(args) == 0 {
		data = msg
	} else {
		data = fmt.Sprintf(msg, args...)
	}
	flash := beego.NewFlash()
	flash.Notice(data)
	flash.Store(c)
}