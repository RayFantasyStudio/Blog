package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"time"
)

type ArticleCreaateController struct  {
	beego.Controller
}
func (c *ArticleCreaateController) Get(){
	c.TplName = "article_create.tpl"
}


func (c *ArticleCreaateController) Post(){
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

	//TODO 处理标签 处理分类

	c.Redirect("/", 302)

}