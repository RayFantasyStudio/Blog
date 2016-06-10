package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"strings"
)

type ArticleCreateController struct {
	beego.Controller
}

func (c *ArticleCreateController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	c.TplName = "article_create.tpl"
}

func (c *ArticleCreateController) Post() {
	article := models.Article{}
	err := c.ParseForm(&article)
	if err != nil {
		beego.Error(err)
	}
	new_category := models.Category{
		Name:article.Category,
	}
	err = models.AddCategory(new_category)
	if err != nil {
		beego.Error(err)
	}
	raw_tags := c.Input().Get("tags")
	tagstr := strings.Split(raw_tags, " ")
	user,err := models.GetUserFromContext(c.Ctx)
	article.Author = user
	err = models.AddArticle(article, tagstr)
	if err != nil {
		beego.Error(err)
	}

	//TODO 处理标签 处理分类

	c.Redirect("/", 302)

}