package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type LoginController  struct {
	beego.Controller
}

func (c *LoginController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		c.Data["Message"] = n
	}
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	var err error
	defer func() {
		if err != nil {
			flash := beego.NewFlash()
			flash.Notice("登陆失败，%s", err)
			flash.Store(&c.Controller)
			c.Redirect(beego.URLFor("LoginController.Get"), 302)
		}
	}()
	user := models.User{}

	if err = c.ParseForm(&user); err != nil {
		beego.Error(err)
		return
	}

	var token string
	token, err = user.Login()

	if err != nil {
		beego.Informational(err)
		return
	}

	c.Ctx.SetCookie("token", token)
	c.Redirect("/", 302)
}

func (c *LoginController) Logout() {
	c.Ctx.SetCookie("token", "")
	c.Redirect("/", 302)
}