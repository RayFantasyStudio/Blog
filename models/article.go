package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id              int64
	Title           string
	Subtitle        string
	Content         string`orm:"size(10000)"`
	Author          string
	Category        string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	ViewCount       int64     `orm:"index"`
	ReplyCount      int64
	LastReplyTime   time.Time `orm:"index"`
	LastReplyUserId int64
}

func GetArticleList(category, tag string) (articles []*Article, err error) {
	o := orm.NewOrm()
	query := o.QueryTable("article")

	if len(category) > 0 {
		query = query.Filter("category", category)
	}
	if len(tag) > 0 {
		query = query.Filter("tag", tag)
	}
	_, err = query.All(&articles)
	return articles, err
}
func GetArticle(id int64) (article *Article, err error) {
	o := orm.NewOrm()
	query := o.QueryTable("article")
	err = query.Filter("id", id).One(article)
	article.ViewCount++
	return article, err
}
func AddAritcle(article *Article) error {
	o := orm.NewOrm()
	_, err := o.Insert(article)
	return err
}