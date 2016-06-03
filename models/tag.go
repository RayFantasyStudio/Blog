package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id   int64
	Name string
}

func AddTag(tag Tag) error {
	o := orm.NewOrm()
	_, err := o.Insert(tag)
	return err
}
func DeleteTag(Id int64, name string) error {
	o := orm.NewOrm()
	tag := new(Tag)
	var err error
	//按照id删除（优先）
	if Id > -1 {
		tag.Id = Id
		_, err = o.Delete(tag)
	}
	//按照名称（name）删除
	if len(name) > 0 {
		tag.Name = name
		_, err = o.Delete(tag)
	}
	return err
}

func GetArticleTag(articleId int64) ([]*Tag, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tag")
	qs.Filter("article_id", articleId)
	tagList := make([]*Tag, 0)
	_, err := qs.All(&tagList)
	return tagList, err
}
func GetTags() ([]Tag, error) {
	o := orm.NewOrm()
	tags := make([]Tag, 0)
	qs := o.QueryTable("tag")
	_, err := qs.All(&tags)
	return tags, err

}