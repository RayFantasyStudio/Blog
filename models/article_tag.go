package models

import (
	"github.com/astaxie/beego/orm"
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

func DeleteArticleTag(articleId int64) (err error){
	_,err = orm.NewOrm().Raw("DELETE FROM article_tag  WHERE article_id = ?",articleId).Exec()
	return
}


func GetArticlesAccroingTagId(tagId int64) (articles []Article,count int64, err error) {
	sql := "SELECT * FROM tag LEFT JOIN article_tag ON tag.id = article_tag.tag_id LEFT JOIN article ON article.id = article_tag.article_id WHERE tag.id= ?"
	count,err = orm.NewOrm().Raw(sql,tagId).QueryRows(&articles)
	return
}
func GetArticlesAccroingTagName(name string) (articles []Article,count int64, err error) {
	sql := "SELECT * FROM tag LEFT JOIN article_tag ON tag.id = article_tag.tag_id LEFT JOIN article ON article.id = article_tag.article_id WHERE tag.name= ?"
	count,err = orm.NewOrm().Raw(sql,name).QueryRows(&articles)
	return
}
func GetTagAccroingAritcle(articleId int64) (tags []Tag,count int64,err error) {
	sql := "SELECT * FROM tag LEFT JOIN article_tag ON tag.id = article_tag.tag_id LEFT JOIN article ON article.id = article_tag.article_id WHERE article_id = ?"
	count,err = orm.NewOrm().Raw(sql,articleId).QueryRows(&tags)
	return
}
