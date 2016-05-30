package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type AdminController  struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplName = "admin.tpl"
}
func (c *AdminController) Article() {
	articleList, err := models.GetArticleList("", "")
	if err != nil {
		beego.Error(err)
	}
	c.Data["articlelist"] = articleList
	c.TplName = "admin_article.tpl"
}
func (c *AdminController) Category() {
	categoryList, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}
	c.Data["categorylist"] = categoryList
	c.TplName = "admin_category.tpl"
}