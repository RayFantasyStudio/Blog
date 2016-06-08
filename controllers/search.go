package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
)

type SearchController  struct {
	beego.Controller
}

func (c *SearchController) Get() {
	key_word := c.Input().Get("key")
	articles,err := models.FindArticles(key_word,true,true,true,true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Articles"] = articles
	c.Data["ResultCount"] = len(articles)
	c.TplName = "search.tpl"
}
func (c *SearchController) Post() {

}