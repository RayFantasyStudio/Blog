package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"time"
)

type ArticleController  struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.TplName = "admin.tpl"
}
func (c *ArticleController) Post() {
	title := c.Input().Get("title")
	subtitle := c.Input().Get("subtitle")
	category := c.Input().Get("category")
	content := c.Input().Get("content")
	article := models.Article{
		Title:title,
		Subtitle:subtitle,
		Category:category,
		Content:content,
		Updated:time.Now(),
		Created:time.Now(),
		LastReplyTime:time.Now(),
	}
	err := models.AddAritcle(&article)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)

}
func (c *ArticleController) Create() {
	c.TplName = "Add_Article.tpl"
}