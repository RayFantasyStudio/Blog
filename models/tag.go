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
func RenameTag(tagId int64,tagName string) error{
	o := orm.NewOrm()
	qs := o.QueryTable("tag").Filter("id",tagId)
	tag := Tag{}
	err := qs.One(&tag)
	if err != nil {
		return err
	}
	tag.Name = tagName
	_,err = o.Update(&tag)
	return err
}
func DeleteTag(tagId int64) error{
	var err error
	o := orm.NewOrm()
	tag := Tag{
		Id:tagId,
	}
	_,err = o.Delete(&tag)
	if err != nil {
		return err
	}

	article_tags := make([]*ArticleTag,0)
	qs := o.QueryTable("article_tag").Filter("tag_id",tagId)
	_,err = qs.All(&article_tags)
	if err != nil {
		return err
	}
	for _,x := range article_tags{
		o.Delete(x)
		if err != nil {
			return err
		}
	}
	return err
}
func FindTags(key string) ([]Tag,error){
	var tags []Tag
	var err error
	o := orm.NewOrm()
	qs := o.QueryTable("tag").Filter("name__iexact",key)
	_,err = qs.All(&tags)
	return  tags,err
}