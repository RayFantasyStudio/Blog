package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"time"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Owner"] = "RayFantasy Studio"
	c.Data["User"] = "访客"
	var articles []*models.Article
	for i := 0; i < 10; i++ {
		article := &models.Article{
			Title:"title" + strconv.Itoa(i),
			Subtitle:"subtitile" + strconv.Itoa(i),
			Content:"content" + strconv.Itoa(i),
			Author:"ztc",
			Created:time.Now(),
			Updated:time.Now(),
		}
		articles = append(articles, article)
	}
	c.Data["Articles"] = articles
	c.TplName = "index.tpl"
}
