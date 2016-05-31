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
func GetArticle(id int64) (*Article, error) {
	o := orm.NewOrm()
	query := o.QueryTable("article")
	article := new(Article)
	err := query.Filter("id", id).One(article)
	article.ViewCount++

	return article, err
}
func AddAritcle(article *Article) error {
	o := orm.NewOrm()
	_, err := o.Insert(article)

	//添加分类
	category := new(Category)
	qs := o.QueryTable("category").Filter("name", article.Category)
	err = qs.One(category)
	if err == nil {
		category.ArticleCount++
		_,err = o.Update(category)
	} else {
		category.Name = article.Category
		category.Created = time.Now()
		category.ArticleCount = 1
		o.Insert(category)
	}
	return err
}
func ModifyArticle(article *Article) error {
	o := orm.NewOrm()
	modifyArticle := &Article{
		Id:article.Id,
	}
	err := o.Read(modifyArticle)
	formerCategory := modifyArticle.Category
	if err == nil {
		modifyArticle.Category = article.Category
		modifyArticle.Title = article.Title
		modifyArticle.Subtitle = article.Subtitle
		modifyArticle.Content = article.Content
		modifyArticle.Updated = time.Now()
	}
	_, err = o.Update(modifyArticle)

	//处理分类
	qs := o.QueryTable("category")
	category := new(Category)
	if len(formerCategory) > 0 {
		err = qs.Filter("name", formerCategory).One(category)
		if err == nil {
			category.ArticleCount--
			_, err = o.Update(category)
		}
	}
	err = qs.Filter("name", article.Category).One(category)
	if err == nil {
		category.ArticleCount++
		o.Update(category)
	} else {
		if len(article.Category) > 0 {
			category.Name = article.Category
			category.Created = time.Now()
			category.ArticleCount++
			o.Insert(category)
		}
	}
	return err
}