package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	var err error
	c.Data["Articles"], err = models.GetArticleList("Created","", "")
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "index.tpl"
}
