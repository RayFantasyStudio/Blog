package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"strconv"
)

type AdminController  struct {
	beego.Controller
}

func (c *AdminController) Get() {

	articles, err := models.GetArticles("", "", "")
	if err != nil {
		beego.Error(err)
	}

	c.Data["Articles"] = articles

	categories, err := models.GetCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories

	replies, err := models.GetAdminReplies()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Replies"] = replies

	tags, err := models.GetTags()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Tags"] = tags

	c.TplName = "admin.tpl"
}
func (c *AdminController) Post() {
	op := c.Input().Get("op")
	beego.Info("op = ", op)
	switch op {
	case "category_rename":
		former_category := c.Input().Get("former_category")
		new_category := c.Input().Get("new_category")
		if (len(former_category) > 0 && len(new_category) > 0) {
			err := models.ModifyCategory(former_category, new_category)
			if err != nil {
				beego.Error(err)
			}
		}
	case "category_delete":
		delete_category_id, err := strconv.ParseInt(c.Input().Get("delete_category_id"), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteCategory(delete_category_id, "")
		if err != nil {
			beego.Error(err)
		}
	}
	c.Redirect("/admin", 302)
}