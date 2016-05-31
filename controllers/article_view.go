package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/RayFantasyStudio/blog/models"
	"time"
)

type ArticleViewController struct {
	beego.Controller
}

func (c *ArticleViewController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	raw_id := c.Input().Get("id")
	id, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article, err := models.GetArticle(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Article"] = article
	c.Data["Replies"], err = models.GetReplies(id)
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "article_view.tpl"
}
func (c *ArticleViewController) Post() {
	raw_aid := c.Input().Get("aid")
	name := c.Input().Get("reply_name")
	content := c.Input().Get("reply_content")
	beego.Error("aid = ",raw_aid,"   name  =  ",name,"   content   =  ",content)
	aid, err := strconv.ParseInt(raw_aid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	reply := models.Reply{
		UserName:name,
		Content:content,
		ArticleId:aid,
		Time:time.Now(),
	}
	err = models.AddReply(&reply)
	c.Redirect("/article/view?id="+raw_aid,302)
}
