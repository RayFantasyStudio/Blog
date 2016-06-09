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
	c.Data["Articles"], err = models.GetArticles(models.Filter_Create, "", "", false)
	if err != nil {
		beego.Error(err)
	}

	page, err := strconv.Atoi(c.Ctx.Input.Param(":page"))
	if err != nil {
		beego.Error(err)
	}

	c.Data["Paginator"] = utils.NewPaginator("/page/%d", page, 10, 200)

	c.TplName = "page.tpl"
}