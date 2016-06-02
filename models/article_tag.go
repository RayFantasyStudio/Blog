package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type ArticleTag  struct {
	Id        int64
	ArticleId int64
	TagId     int64
}

func SaveArticleTag(article_id, tag_id int64) error{
	o := orm.NewOrm()
	article_tag := new(ArticleTag)
	article_tag.ArticleId = article_id
	article_tag.TagId = tag_id
	_,err := o.Insert(&article_tag)
	return err
}

func GetArticleAccroingTag(tag_id int64) ([]*Article, error) {
	o := orm.NewOrm()
	qs_tag := o.QueryTable("article_tag").Filter("tag_id", tag_id)
	m2m_article := make([]int64, 0)
	_, err := qs_tag.All(m2m_article)


	qs_article := o.QueryTable("article")
	articles := make([]*Article, 0)
	for _, x := range m2m_article {
		qs_article.Filter("id", x)
		article := new(Article)
		qs_article.One(article)
		articles = append(articles, article)
	}
	return articles, err
}
func GetTagAccroingAritcle(article_id int64) ([]Tag, error) {
	o := orm.NewOrm()
	qs_article := o.QueryTable("article_tag").Filter("article_id", article_id)
	m2m_tags := make([]ArticleTag, 0)
	_, err := qs_article.All(&m2m_tags)



	//FIXME 一下功能未实现，从tag表取出标签
	beego.Info(m2m_tags)
	qs_tag := o.QueryTable("tag")
	tags := make([]Tag, 0)
	for _, x := range m2m_tags {
		qs_tag.Filter("id", x.TagId)
		var tag Tag
		qs_tag.One(&tag)
		tags = append(tags, tag)
	}
	return tags, err
}
