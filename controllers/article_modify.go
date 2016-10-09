package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"strconv"
	"strings"
)

type ArticleModifyController struct {
	beego.Controller
}

func (c *ArticleModifyController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	raw_id := c.Input().Get("id")
	beego.Info("the raw id =  ", raw_id)
	id, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article, err := models.GetArticle(id)
	c.Data["Article"] = article

	tags,_,err := models.GetTagAccroingAritcle(id)
	var tag_string []string
	for _,x := range tags{
		tag_string = append(tag_string,x.Name)
	}
	tag_str := strings.Join(tag_string," ")
	c.Data["Tags"] = tag_str
	c.TplName = "article_modify.tpl"
}

func (c *ArticleModifyController) Post() {

	eUser,err := models.GetUserFromContext(c.Ctx)
	if err != nil {
		beego.Error(err)
	}
	raw_id := c.Input().Get("id")
	former_tags := c.Input().Get("former_tags")
	tags := c.Input().Get("tags")
	id, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article := models.Article{}
	err = c.ParseForm(&article)
	if err != nil {
		beego.Error(err)
	}
	article.Author = eUser
	article.Id = id
	err = models.ModifyArticle(&article,former_tags,tags)
	if err != nil {
		beego.Error(err)
	}

	//TODO 处理标签 处理分类

	c.Redirect("/", 302)

}
