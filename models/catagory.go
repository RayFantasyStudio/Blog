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

func GetCategoryList() ([]*Category, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	categoryList := make([]*Category, 0)
	_, err := qs.All(&categoryList)
	return categoryList, err
}
func AddCagegory(category Category) error {
	o := orm.NewOrm()
	_, err := o.Insert(category)
	return err
}
func DeleteCategory(Id int64, name string) error {
	o := orm.NewOrm()
	category := new(Category)
	var err error
	//按照id删除（优先）
	if Id > -1 {
		category.Id = Id
		_,err = o.Delete(category)
	}
	//按照名称（name）删除
	if len(name) > 0{
		category.Name = name
		_,err = o.Delete(name)
	}
	return err
}