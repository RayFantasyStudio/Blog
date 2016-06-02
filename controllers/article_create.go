package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"time"
	"strings"
)

type ArticleCreaateController struct  {
	beego.Controller
}
func (c *ArticleCreaateController) Get(){
	initHeaderFooterData(&c.Controller, owner + "的Blog")
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
	raw_tags := c.Input().Get("tags")
	tagstr := strings.Split(raw_tags," ")
	err := models.AddAritcle(&article,tagstr)
	if err != nil {
		beego.Error(err)
	}

	//TODO 处理标签 处理分类

	c.Redirect("/", 302)

}