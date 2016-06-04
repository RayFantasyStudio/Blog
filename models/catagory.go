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

func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	categoryList := make([]*Category, 0)
	_, err := qs.All(&categoryList)
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
	_,err = qs_article.All(articles)

	for _,x := range articles{
		x.Category = category_name
		_,err = o.Update(x)
	}
	return err
}
