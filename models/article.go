package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Article struct {
	Id              int64
	Title           string
	Subtitle        string
	Content         string`orm:"size(10000)"`
	Author          *User `orm:"rel(fk)"`
	Category        string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	ViewCount       int64     `orm:"index"`
	ReplyCount      int64
	LastReplyTime   time.Time `orm:"index"`
	LastReplyUserId int64
}

func GetArticles(order_key string, category, tag string) (articles []*Article, err error) {
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
	article := Article{}
	err := query.Filter("id", id).One(&article)
	if err != nil {
		return nil, err
	}
	article.ViewCount++
	o.Update(&article)
	return &article, err
}
func AddArticle(article Article, tagsrt []string) error {
	beego.Info(tagsrt)
	o := orm.NewOrm()
	var err error
	//插入文章
	_, err = o.Insert(&article)
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
			_, err = o.Update(&category)
			if err != nil {
				return err
			}
		}
	}


	//添加Tag
	var inexistTags []int64
	for _, x := range tagsrt {
		tag := Tag{Name:x}
		_, _, err := o.ReadOrCreate(&tag, "Name")
		if err != nil {
			return err
		} else {
			inexistTags = append(inexistTags, tag.Id)
		}
	}
	article_tmp := Article{}
	qs_art := o.QueryTable("article").Filter("title", article.Title)
	err = qs_art.One(&article_tmp)
	if err != nil {
		return err
	}

	for _, x := range inexistTags {
		err = SaveArticleTag(article_tmp.Id, x)
		if err != nil {
			return err
		}
	}

	return err
}
func ModifyArticle(article *Article) error {
	o := orm.NewOrm()
	modifyArticle := Article{
		Id:article.Id,
	}
	err := o.Read(&modifyArticle)
	formerCategory := modifyArticle.Category
	if err == nil {
		modifyArticle.Category = article.Category
		modifyArticle.Title = article.Title
		modifyArticle.Subtitle = article.Subtitle
		modifyArticle.Content = article.Content
		modifyArticle.Updated = time.Now()
	}
	_, err = o.Update(&modifyArticle)
	if err != nil {
		return err
	}

	//处理分类
	qs := o.QueryTable("category")
	category := Category{}
	if len(formerCategory) > 0 {
		err = qs.Filter("name", formerCategory).One(&category)
		if err == nil {
			category.ArticleCount--
			_, err = o.Update(&category)
		}
	}
	err = qs.Filter("name", article.Category).One(&category)
	if err == nil {
		category.ArticleCount++
		o.Update(category)
	} else {
		if len(article.Category) > 0 {
			category.Name = article.Category
			category.Created = time.Now()
			category.ArticleCount++
			o.Insert(&category)
		}
	}
	return err
}
func DeleteArticle(id int64) error {
	o := orm.NewOrm()
	cate := Article{Id : id}
	_, err := o.Delete(&cate)
	return err
}