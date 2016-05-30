package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type AdminController  struct{
	beego.Controller
}

func (c *AdminController) Get()  {
	articleList,err := models.GetArticleList("","")
	if err != nil {
		beego.Error(err)
	}
	c.Data["articlelist"] = articleList
	c.TplName = "admin.tpl"
}