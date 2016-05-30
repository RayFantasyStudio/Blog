package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type AdminController  struct {
	beego.Controller
}

func (c *AdminController) Get() {
	categories, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories

	articles, err := models.GetArticleList("", "")
	if err != nil {
		beego.Error(err)
	}
	c.Data["Articles"] = articles

	c.TplName = "admin.tpl"
}