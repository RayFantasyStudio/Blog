package controllers

import (
	"github.com/astaxie/beego"

	"strconv"
	"github.com/RayFantasyStudio/blog/models"
	"github.com/RayFantasyStudio/blog/utils"
)

type ArticleController  struct {
	beego.Controller
}
type articleItem struct{
	article models.Article
	tag	[]models.Tag
	authorName string
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
	page,err := c.GetInt("page")
	if err != nil {
		beego.Error(err)
	}
	order_by := c.GetString("order")
	desc,err := c.GetBool("desc")
	if err != nil {
		beego.Error(err)
	}
	beego.Info("the page is ",page,"the order is ",order_by,"   ",desc)
	c.Data["Paginator"] = utils.NewPaginator("/article?page=%d&order=create&desc=true", page, models.ArticlePerPageLimit, models.GetTotalArticleCount())
	var articles []*models.Article
	articles,_,err = models.GetArticles(order_by,"","",desc,page)
	c.Data["Articles"] = articles
	author_name := make([]string,len(articles))
	for index,x := range articles{
		author_name[index],err = models.GetUserNameById(x.Author.Id)
		if err != nil {
			beego.Error(err)
		}
	}
	beego.Info(author_name)
	c.Data["Authors"] = author_name
	c.TplName = "article.tpl"
}
func (c *ArticleController) Post() {

}
