package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"strconv"
)

type ArticleModifyController struct {
	beego.Controller
}

func (c *ArticleModifyController) Get() {
	raw_id := c.Input().Get("id")
	beego.Info("the raw id =  ", raw_id)
	id, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article, err := models.GetArticle(id)
	c.Data["Article"] = article
	c.TplName = "article_modify.tpl"
}

func (c *ArticleModifyController) Post() {
	raw_id := c.Input().Get("id")
	title := c.Input().Get("title")
	subtitle := c.Input().Get("subtitle")
	category := c.Input().Get("category")
	content := c.Input().Get("content")
	id, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article := models.Article{
		Id:id,
		Title:title,
		Subtitle:subtitle,
		Category:category,
		Content:content,
	}
	err = models.ModifyArticle(&article)
	if err != nil {
		beego.Error(err)
	}
	//TODO 处理标签 处理分类

	c.Redirect("/", 302)

}
