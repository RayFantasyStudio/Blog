package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"time"
	"strings"
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

	//处理标签
	raw_tags := c.Input().Get("tags")
	tags := strings.Split(raw_tags," ")
	for _,x := range tags{
		tag := models.Tag{
			Name:x,
		}
		err := models.AddTag(tag)
		if err != nil {
			beego.Error(err)
			break;
		}
	}
	c.Redirect("/", 302)

}
func (c *ArticleController) Create() {
	c.TplName = "Add_Article.tpl"
}
func (c *ArticleController) View(){
	c.TplName = "Article_View.tpl"
}