package controllers

import (
	"github.com/astaxie/beego"
	"github.com/RayFantasyStudio/blog/models"
	"strconv"
	"path"
	"path/filepath"
)

//admin操作op
const (
	article_delete = "article_delete"
	category_rename = "category_rename"
	category_delete = "category_delete"
	tag_rename = "tag_rename"
	tag_delete = "tag_delete"
	reply_delete = "reply_delete"

	account_setting = "account_setting"

	article_filter = "article_filter"
	category_filter = "category_filter"
)

type AdminController  struct {
	beego.Controller
}

func (c *AdminController) Get() {
	//处理文章过滤器
	var articles []*models.Article
	var err error
	article_category_filter := c.GetSession("article_category_filter")
	if article_category_filter == nil || article_category_filter == "All" {
		article_category_filter = ""
	}
	article_order_filter := c.GetSession("article_order")
	beego.Info(article_order_filter)
	switch article_order_filter {
	case "create_ascending":
		articles, _, err = models.GetArticles(models.Filter_Create, article_category_filter.(string), "", int64(0), true, 1)
	case "create_descending":
		articles, _, err = models.GetArticles(models.Filter_Create, article_category_filter.(string), "", int64(0), false, 1)
	case "update_ascending":
		articles, _, err = models.GetArticles(models.Filter_Update, article_category_filter.(string), "", int64(0), true, 1)
	case "update_descending":
		articles, _, err = models.GetArticles(models.Filter_Update, article_category_filter.(string), "", int64(0), false, 1)
	case "last_reply_time_ascending":
		articles, _, err = models.GetArticles(models.Filter_LastReplyTime, article_category_filter.(string), "", int64(0), true, 1)
	case "last_reply_time_descending":
		articles, _, err = models.GetArticles(models.Filter_LastReplyTime, article_category_filter.(string), "", int64(0), false, 1)
	case "view_count_ascending":
		articles, _, err = models.GetArticles(models.Filter_ViewCount, article_category_filter.(string), "", int64(0), true, 1)
	case "view_count_descending":
		articles, _, err = models.GetArticles(models.Filter_ViewCount, article_category_filter.(string), "", int64(0), false, 1)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Data["Articles"] = articles

	//处理分类过滤器
	var categories []*models.Category
	category_order_filter := c.GetSession("category_order")
	switch category_order_filter {
	case "category_create_ascending":
		categories, err = models.GetCategories(models.Filter_Category_Create, false)
	case "category_create_descending":
		categories, err = models.GetCategories(models.Filter_Category_Create, true)
	case "category_article_count_ascending":
		categories, err = models.GetCategories(models.Filter_Category_ArticleCount, false)
	case "category_article_count_descending":
		categories, err = models.GetCategories(models.Filter_Category_ArticleCount, true)
	}
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

	user, err := models.GetUserFromContext(c.Ctx)
	if err != nil {
		beego.Error(err)
	}
	c.Data["UserName"] = user.Name

	var avatar_path string
	avatar_path,err = models.GetUserAvatar(int64(user.Id))
	c.Data["avatar_path"] = avatar_path

	c.TplName = "admin.tpl"
}
func (c *AdminController) Post() {
	op := c.Input().Get("op")
	beego.Info("op = ", op)
	switch op {
	case article_delete:
		aid, err := strconv.ParseInt(c.Input().Get("delete_article_id"), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteArticle(aid)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin#/article", 302)
	case category_rename:
		former_category := c.Input().Get("former_category")
		new_category := c.Input().Get("new_category")
		if (len(former_category) > 0 && len(new_category) > 0) {
			err := models.ModifyCategory(former_category, new_category)
			if err != nil {
				beego.Error(err)
			}
		}
		c.Redirect("/admin#/category", 302)
	case category_delete:
		delete_category_id, err := strconv.ParseInt(c.Input().Get("delete_category_id"), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteCategory(delete_category_id, "")
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin#/category", 302)
	case tag_rename:
		rename_tag_id, err := strconv.ParseInt(c.Input().Get("rename_tag_id"), 10, 64)
		rename_tag_name := c.Input().Get("rename_tag_name")
		beego.Info("rename tag id = ", rename_tag_id)
		if err != nil {
			beego.Error(err)
		}
		if ( len(rename_tag_name) > 0) {
			err := models.RenameTag(rename_tag_id, rename_tag_name)
			if err != nil {
				beego.Error(err)
			}
		}
		c.Redirect("/admin#/tag", 302)
	case tag_delete:
		delete_tag_id, err := strconv.ParseInt(c.Input().Get("delete_tag_id"), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteTag(delete_tag_id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin#/tag", 302)
	case reply_delete:
		delete_reply_id, err := strconv.ParseInt(c.Input().Get("delete_reply_id"), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteReply(delete_reply_id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin#/reply", 302)
	case article_filter:
		c.SetSession("article_category_filter", c.Input().Get("article_category_filter"))
		c.SetSession("article_order", c.Input().Get("order"))
		c.Redirect("/admin#/article", 302)
	case category_filter:
		c.SetSession("category_order", c.Input().Get("category_order"))
		c.Redirect("/admin", 302)
	case account_setting:
		user_name := c.Input().Get("user_name")
		_, fh, err := c.GetFile("avatar")
		if err != nil {
			beego.Error(err)
		}
		beego.Info("Saving................")
		var attachment string
		if fh != nil {
			extension := filepath.Ext(fh.Filename)
			fh.Filename = user_name + "_avatar" + extension
			attachment = fh.Filename
			path := path.Join("static", "img", "user_avatar", attachment)
			err = c.SaveToFile("avatar", path )
			if err != nil {
				beego.Error(err)
			}
			user,err := models.GetUserFromContext(c.Ctx)
			if err != nil {
				beego.Error(err)
			}
			beego.Info("uid = ",user.Id,"  path = ",path)
			err = models.SaveUserAvatar(int64(user.Id),path)
			if err != nil {
				beego.Error(err)
			}
		}
		//处理用户名、密码

		user,err := models.GetUserFromContext(c.Ctx)
		if err != nil {
			beego.Error(err)
		}
		uid := int64(user.Id)
		username :=  c.Input().Get("username")
		pwd := c.Input().Get("pwd")
		beego.Info("The current user has been change \n username : ",username,"\n pwd :",pwd,"\nid : ",uid)
		err = models.ModifyUserInfo(uid,username,pwd)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Redirect("/admin", 302)
}