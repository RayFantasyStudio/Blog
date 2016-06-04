package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
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

func GetArticleList(order_key string, category, tag string) (articles []*Article, err error) {
	o := orm.NewOrm()
	query := o.QueryTable("article")
	switch order_key {
	case "Created":
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		_, err = query.OrderBy("-created").All(&articles)
		return articles, err
	case "LastReplyTime":
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		_, err = query.OrderBy("-last_reply_time").All(&articles)
		return articles, err
	case "ViewCount":
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		_, err = query.OrderBy("-view_count").All(&articles)
		return articles, err
	}

	_, err = query.All(&articles)
	return articles, err
}
func GetArticle(id int64) (*Article, error) {
	o := orm.NewOrm()
	query := o.QueryTable("article")
	article := new(Article)
	err := query.Filter("id", id).One(article)
	if err != nil {
		return nil,err
	}
	article.ViewCount++
	_,err = o.Update(&article)
	return article, err
}
func AddArticle(article *Article, tagsrt []string) error {
	beego.Info(tagsrt)
	o := orm.NewOrm()
	var err error
	//插入文章
	_,err = o.Insert(article)
	if err != nil {
		return err
	}
	//添加分类
	category := Category{Name:article.Category}
	if created, _, err := o.ReadOrCreate(&category, "Name"); err == nil {
		if created {
			category.ArticleCount = 1
		} else {
			category.ArticleCount++
			o.Update(&category)
		}
	}

	
	//添加Tag
	tags_id := make([]int64, 0)
	for _, x := range tagsrt {
		tag := Tag{Name:x}
		if created, tag_id, err := o.ReadOrCreate(&tag, "Name"); err == nil {
			if created {
				fmt.Println("New Insert an object. Id:", tag.Name)
				tags_id = append(tags_id, tag_id)
			} else {
				fmt.Println("Get an object. Id:", tag.Name)
			}
		}
	}
	//处理article - tag
	article_tmp := new(Article)
	qs_art := o.QueryTable("article")
	err = qs_art.Filter("Title", article.Title).One(article_tmp)
	for _, x := range tags_id {
		SaveArticleTag(article_tmp.Id, x)

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
func DeleteArticle(id int64) error {
	o := orm.NewOrm()
	cate := &Article{Id : id}
	_, err := o.Delete(cate)

	return err
}