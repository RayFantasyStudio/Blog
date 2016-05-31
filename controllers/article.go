package controllers

import (
	"github.com/astaxie/beego"

	"strconv"
	"github.com/RayFantasyStudio/blog/models"
)

type ArticleController  struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	op := c.Input().Get("op")
	if op == "del" {
		id, err := strconv.ParseInt(c.Input().Get("id"), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteArticle(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/",302)
	}
	c.TplName = "admin.tpl"
}
func (c *ArticleController) Post() {

}
