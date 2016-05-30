package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"

	"blog_allen/controllers/ExtraUnit"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Owner"] = "RayFantasy Studio"
	c.Data["User"] = "访客"
	var err error
	c.Data["Articles"],err  = models.GetArticleList("","")
	if err != nil {
		beego.Error(err)
	}
	beego.AddFuncMap("SinceTime",controllers.SinceTime)

	c.TplName = "index.tpl"
}
