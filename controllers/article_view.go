package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/RayFantasyStudio/blog/models"
)

type ArticleViewController struct {
	beego.Controller
}

func (c *ArticleViewController) Get() {
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
	c.TplName = "article_view.tpl"
}
