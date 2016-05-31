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
	c.TplName = "article_view.tpl"
	raw_id := c.Ctx.Input.Param("0")
	beego.Info("the raw id =  ", raw_id)
	_, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article, err := models.GetArticle(2)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Article"] = article
}
