package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strings"
)
//文章字段过滤器
const (
	Filter_Create = "created"
	Filter_Update = "updated"
	Filter_ViewCount = "view_count"
	Filter_LastReplyTime = "last_reply_time"
)

type Article struct {
	Id              int64
	Title           string
	Subtitle        string
	Content         string`orm:"size(10000)"`
	Author          *User`orm:"rel(fk)"`
	Category        string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	ViewCount       int64     `orm:"index"`
	ReplyCount      int64
	LastReplyTime   time.Time `orm:"index"`
	LastReplyUserId int64
}

func GetArticles(order_key string, category, tag string, inverted bool) (articles []*Article, err error) {
	o := orm.NewOrm()
	query := o.QueryTable("article")
	switch order_key {
	case Filter_Create:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_Create).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_Create).All(&articles)
		}
		return articles, err
	case Filter_LastReplyTime:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_LastReplyTime).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_LastReplyTime).All(&articles)
		}
		return articles, err
	case Filter_ViewCount:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_ViewCount).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_ViewCount).All(&articles)
		}
		return articles, err
	case Filter_Update:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_Update).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_Update).All(&articles)
		}
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
func ModifyArticle(article *Article, raw_former_tag string, raw_tags string) error {
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
		modifyArticle.Author = article.Author
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
		o.Update(&category)
	} else {
		if len(article.Category) > 0 {
			category.Name = article.Category
			category.Created = time.Now()
			category.ArticleCount++
			o.Insert(&category)
		}
	}

	//FIXME 标签处理仍存在问题
	//处理标签
	former_tag := strings.Split(raw_former_tag, " ")
	tags := strings.Split(raw_tags, " ")
	var deleted_tag, created_tag []string
	var flag bool = false
	for x, _ := range former_tag {
		for y, _ := range tags {
			if former_tag[x] == tags[y] {
				flag = true
				break
			}
		}
		if !flag {
			deleted_tag = append(deleted_tag, former_tag[x])
			flag = false
		}
	}

	for _, x := range tags {
		for a, y := range former_tag {
			if y == x {
				break
			}
			if a == len(tags) && y != x {
				deleted_tag = append(created_tag, x)
			}
		}
	}
	beego.Info("the create list of ", created_tag)
	beego.Info("the delete list of ", deleted_tag)
	beego.Info("the raw create list of ", former_tag)
	beego.Info("the raw delete list of ", tags)
	return err
}
func DeleteArticle(id int64) error {
	o := orm.NewOrm()
	cate := Article{Id : id}
	_, err := o.Delete(&cate)
	if err != nil {
		return err
	}
	var delete_list []ArticleTag
	qs := o.QueryTable("article_tag").Filter("article_id",id)
	_,err = qs.All(&delete_list)
	if err != nil {
		return err
	}
	for _,x := range delete_list {
		_,err = o.Delete(&x)
		if err != nil {
			return err
		}
	}
	return err
}
func FindArticles(key string,byTitle bool,bySubtitle bool,byCategory bool,byTag bool) ([]Article,error){
	o := orm.NewOrm()
	var articles []Article
	var err error
	//按照标题检索
	if byTitle{
		qs_article_title := o.QueryTable("article").Filter("title__iexact", key)
		var articles_tmp []Article
		_,err = qs_article_title.All(&articles_tmp)
		for _,x:= range articles_tmp{
			articles = append(articles,x)
		}
	}

	//按照副标题检索
	if byTitle{
		qs_article_title := o.QueryTable("article").Filter("subtitle__iexact", key)
		var articles_tmp []Article
		_,err = qs_article_title.All(&articles_tmp)
		for _,x:= range articles_tmp{
			articles = append(articles,x)
		}
	}

	//按照分类检索
	if byTitle{
		qs_article_title := o.QueryTable("article").Filter("category__iexact", key)
		var articles_tmp []Article
		_,err = qs_article_title.All(&articles_tmp)
		for _,x:= range articles_tmp{
			articles = append(articles,x)
		}
	}

	//按照标签检索
	if byTag{
		var article_tmp Article
		var tags_tmp []Tag
		var article_tags_tmp []ArticleTag
		qs_tag := o.QueryTable("article")
		tags_tmp,err = FindTags(key)
		if err != nil {
			return nil,err
		}
		article_tags_tmp,err = FindArticleTagFromTags(tags_tmp)
		if err != nil {
			return nil,err
		}

		for _,x := range article_tags_tmp{
			qs_tag = qs_tag.Filter("id",x.ArticleId)
			err = qs_tag.One(&article_tmp)
			if err != nil {
				return nil,err
			}
			articles = append(articles,article_tmp)
		}
	}
	return articles,err
}