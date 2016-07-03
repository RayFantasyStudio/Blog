package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"github.com/RayFantasyStudio/blog/utils"
	"strconv"
)

type PageController struct {
	beego.Controller
}

func (c *PageController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	var err error
	var articles []*models.Article

	page, err := strconv.Atoi(c.Ctx.Input.Param(":page"))
	if err != nil {
		beego.Error(err)
	}

	c.Data["Paginator"] = utils.NewPaginator("/page/%d", page, models.ArticlePerPageLimit, models.GetTotalArticleCount())
	articles, _, err = models.GetArticles(models.Filter_Create, "", "",0, true, page)
	c.Data["Articles"] = articles
	if err != nil {
		beego.Error(err)
	}
	author_name := make([]string,len(articles))
	for index,x := range articles{
		author_name[index],err = models.GetUserNameById(x.Author.Id)
		if err != nil {
			beego.Error(err)
		}
	}
	beego.Info(author_name)
	c.Data["Authors"] = author_name
	c.TplName = "page.tpl"
}