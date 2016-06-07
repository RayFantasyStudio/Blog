package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"github.com/RayFantasyStudio/blog/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	var err error
	c.Data["Articles"], err = models.GetArticles(models.Filter_Create, "", "", false)
	if err != nil {
		beego.Error(err)
	}

	// FIXME: 测试结束后改为路由分发形式，下同
	page, _ := c.GetInt("page")

	c.Data["Paginator"] = utils.NewPaginator(c.Ctx.Request, page, 10, 200)

	c.TplName = "index.tpl"
}
