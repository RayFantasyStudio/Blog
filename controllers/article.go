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
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	var article_count int
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
	uid,err := c.GetInt64("by_uid")
	if err != nil {
		beego.Error(err)
	}
	category := c.GetString("cate")

	c.Data["Category_tmp"] = category
	c.Data["Order_tmp"] = order_by
	c.Data["User_tmp"] = uid
	c.Data["Desc_tmp"] = desc





	beego.Info("the page is ",page,"the order is ",order_by,"   ","desc = ",desc,"cate = ",category)
	var articles []*models.Article
	articles, article_count,err = models.GetArticles(order_by,category,"",uid,desc,page)
	if err != nil {
		beego.Error(err)
	}
	beego.Info("the result of get articles = ",article_count)
	c.Data["Articles"] = articles

	c.Data["Paginator"] = utils.NewPaginator("/article?page=%d&order=created&desc=true&by_uid="+strconv.FormatInt(uid,10), page, models.ArticlePerPageLimit, int64(article_count))

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
