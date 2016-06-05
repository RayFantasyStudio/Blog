package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id           int64
	Name         string
	Created      time.Time`orm:"index"`
	ArticleCount int64`orm:"index"`
}
//分类字段过滤器
const (
	Filter_Category_Create = "created"
	Filter_Category_ArticleCount = "article_count"
)
func GetCategories(order_key string,inverted bool) ([]*Category, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	categoryList := make([]*Category, 0)
	var err error
	switch order_key {
	case Filter_Category_Create:
		if inverted{
			_, err = qs.OrderBy("-" + Filter_Category_Create).All(&categoryList)
		}else {
			_, err = qs.OrderBy(Filter_Category_Create).All(&categoryList)
		}
	case Filter_Category_ArticleCount:
		if inverted{
			_, err = qs.OrderBy("-" + Filter_Category_ArticleCount).All(&categoryList)
		}else {
			_, err = qs.OrderBy(Filter_Category_ArticleCount).All(&categoryList)
		}
	}
	return categoryList, err
}
func AddCategory(category Category) error {
	o := orm.NewOrm()
	category.Created = time.Now()
	_, err := o.Insert(&category)
	return err
}
func DeleteCategory(Id int64, name string) error {
	o := orm.NewOrm()
	category := new(Category)
	var err error
	//按照id删除（优先）
	if Id > -1 {
		category.Id = Id
		_, err = o.Delete(category)
		if err != nil {
			return err
		}
	}
	//按照名称（name）删除
	if len(name) > 0 {
		category.Name = name
		_, err = o.Delete(name)
		if err != nil {
			return err
		}
	}
	return err
}
func ModifyCategory(former_category, category_name string) error {
	o := orm.NewOrm()
	qs := o.QueryTable("category").Filter("name",former_category)
	category_tmp := Category{}
	err := qs.One(&category_tmp)
	category_tmp.Name = category_name
	category_tmp.Created = time.Now()
	_,err = o.Update(&category_tmp)

	articles := make([]Article,0)
	qs_article := o.QueryTable("article").Filter("category",former_category)
	_,err = qs_article.All(&articles)

	for _,x := range articles{
		x.Category = category_name
		_,err = o.Update(&x)
	}
	return err
}
