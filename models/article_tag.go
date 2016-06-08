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
	article_tag := ArticleTag{
		ArticleId:article_id,
		TagId:tag_id,
	}
	_,err := o.Insert(&article_tag)
	return err
}

func DeleteArticleTag(article_id, tag_id int64) error{
	o := orm.NewOrm()
	article_tag := ArticleTag{
		ArticleId:article_id,
		TagId:tag_id,
	}
	_,err := o.Delete(&article_tag)
	return err
}

func GetArticleAccroingTag(tag_id int64) ([]Article, error) {
	o := orm.NewOrm()
	qs_tag := o.QueryTable("article_tag").Filter("tag_id", tag_id)
	m2m_article := make([]ArticleTag, 0)
	_, err := qs_tag.All(&m2m_article)


	qs_article := o.QueryTable("article")
	articles := make([]Article, 0)
	for _, x := range m2m_article {
		qs_article := qs_article.Filter("id", x.ArticleId)
		article := Article{}
		qs_article.One(&article)
		articles = append(articles, article)
	}
	return articles, err
}
func GetTagAccroingAritcle(article_id int64) ([]Tag, error) {
	o := orm.NewOrm()
	qs_article := o.QueryTable("article_tag").Filter("article_id", article_id)
	m2m_tags := make([]ArticleTag, 0)
	_, err := qs_article.All(&m2m_tags)



	beego.Info(m2m_tags,"len",len(m2m_tags))
	qs_tag := o.QueryTable("tag")
	tags := make([]Tag, len(m2m_tags))
	for c, x := range m2m_tags {
		qs_tag := qs_tag.Filter("id", x.TagId)
		qs_tag.One(&tags[c])
	}
	return tags, err
}
func FindArticleTagFromTags(tags []Tag) ([]ArticleTag,error){
	var article_tags []ArticleTag
	var err error
	o := orm.NewOrm()
	qs := o.QueryTable("article_tag")
	for _,x := range tags {
		qs = qs.Filter("tag_id",x.Id)
		var article_tag_tmp ArticleTag
		err = qs.One(&article_tag_tmp)
		if err != nil {
			return nil,err
		}
		article_tags = append(article_tags,article_tag_tmp)
	}
	return article_tags,err
}
