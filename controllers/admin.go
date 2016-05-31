package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type AdminController  struct {
	beego.Controller
}

func (c *AdminController) Get() {

	articles, err := models.GetArticleList("", "")
	if err != nil {
		beego.Error(err)
	}
	c.Data["Articles"] = articles
	categories, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories


	replies, err := models.GetReplies()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Replies"] = replies


	c.TplName = "admin.tpl"
}