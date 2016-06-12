package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/RayFantasyStudio/blog/models"
	"time"
)

type ArticleViewController struct {
	beego.Controller
}

func (c *ArticleViewController) Get() {
	initHeaderFooterData(&c.Controller, owner + "的Blog")
	raw_id := c.Input().Get("id")
	id, err := strconv.ParseInt(raw_id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article, err := models.GetArticle(id)
	if err != nil {
		beego.Error(err)
	}
	var replies []*models.Reply
	c.Data["Article"] = article
	replies, err = models.GetReplies(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Replies"] = replies
	//加载Tag
	tags, err := models.GetTagAccroingAritcle(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Tags"] = tags


	//引用回复
	quote_id, err := c.GetInt64("quote_rid")
	if err != nil {
		beego.Error(err)
	}
	c.Data["quote_rid"] = quote_id
	if quote_id > 0 {
		quote_reply, err := models.GetReplyById(quote_id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["quote_reply"] = quote_reply
	}
	quote_map := make(map[int64]models.Reply)
	//显示引用回复
	reply_tmp := models.Reply{}
	for _, x := range replies {
		if x.QuoteReplyId > 0 {
			reply_tmp, err = models.GetReplyById(x.QuoteReplyId)
			quote_map[x.QuoteReplyId] = reply_tmp
		}
	}
	beego.Info(quote_map)
	c.Data["quote_map"] = quote_map

	//删除评论
	if c.Input().Get("op") == "delrpy" {
		raw_rid := c.Input().Get("rid")
		raw_aid := c.Input().Get("aid")
		rid, err := strconv.ParseInt(raw_rid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteReply(rid)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/article/view?id=" + raw_aid, 302)
	}
	c.TplName = "article_view.tpl"
}
func (c *ArticleViewController) Post() {
	raw_aid := c.Input().Get("aid")
	name := c.Input().Get("reply_name")
	content := c.Input().Get("reply_content")
	quote_rid, err := c.GetInt64("quote_reply_id")
	if err != nil {
		beego.Error(err)
	}
	aid, err := strconv.ParseInt(raw_aid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	reply := models.Reply{
		UserName:name,
		Content:content,
		ArticleId:aid,
		Time:time.Now(),
	}
	if quote_rid > 0 {
		reply.QuoteReplyId = quote_rid
	}

	err = models.AddReply(&reply)
	c.Redirect("/article/view?id=" + raw_aid, 302)
}
