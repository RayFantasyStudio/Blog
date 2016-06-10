package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strings"
	"fmt"
)
//文章字段过滤器
const (
	Filter_Create = "created"
	Filter_Update = "updated"
	Filter_ViewCount = "view_count"
	Filter_LastReplyTime = "last_reply_time"

	ArticlePerPageLimit = 10
)

var totalArticleCount int64

type Article struct {
	Id              int64
	Title           string `form:"title"`
	Subtitle        string `form:"subtitle"`
	Content         string `orm:"size(10000)" form:"content"`
	Author          *User  `orm:"rel(fk)"`
	Category        string `form:"category"`
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	ViewCount       int64     `orm:"index"`
	ReplyCount      int64
	LastReplyTime   time.Time `orm:"index"`
	LastReplyUserId int64
}

func GetArticles(order_key string, category, tag string, inverted bool, page int) (articles []*Article, article_count int, err error) {
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
			_, err = query.OrderBy("-" + Filter_Create).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_Create).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		}
		article_count = len(articles)
		return articles, article_count, err
	case Filter_LastReplyTime:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_LastReplyTime).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_LastReplyTime).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		}
		article_count = len(articles)
		return articles, article_count, err
	case Filter_ViewCount:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_ViewCount).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_ViewCount).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		}
		article_count = len(articles)
		return articles, article_count, err
	case Filter_Update:
		if len(category) > 0 {
			query = query.Filter("category", category)
		}
		if len(tag) > 0 {
			query = query.Filter("tag", tag)
		}
		if inverted {
			_, err = query.OrderBy("-" + Filter_Update).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		} else {
			_, err = query.OrderBy(Filter_Update).Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
		}
		article_count = len(articles)
		return articles, article_count, err
	}
	_, err = query.Limit(ArticlePerPageLimit).Offset((page - 1) * ArticlePerPageLimit).All(&articles)
	article_count = len(articles)
	return articles, article_count, err
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
	qs := o.QueryTable("article_tag").Filter("article_id", id)
	_, err = qs.All(&delete_list)
	if err != nil {
		return err
	}
	for _, x := range delete_list {
		_, err = o.Delete(&x)
		if err != nil {
			return err
		}
	}
	return err
}

func FindArticles(key string, byTitle bool, bySubtitle bool, byCategory bool, byTag bool) (articles []*Article, err error) {
	var conds []string
	baseCond := fmt.Sprintf("article.%%s like '%%%%%s%%%%'", key)
	// 按照标题检索
	if byTitle {
		conds = append(conds, fmt.Sprintf(baseCond, "title"))
	}

	// 按照副标题检索
	if bySubtitle {
		conds = append(conds, fmt.Sprintf(baseCond, "subtitle"))
	}

	// 按照分类检索
	if byCategory {
		conds = append(conds, fmt.Sprintf(baseCond, "category"))
	}

	if byTag {
		conds = append(conds, fmt.Sprintf("tag.name like '%%%s%%'", key))
	}

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}

	// 构造SQL语句
	qb.Select("article.*").From("tag").
	InnerJoin("(article INNER JOIN article_tag ON article.id = article_tag.article_id)").
	On("tag.id = article_tag.tag_id")
	if len(conds) > 0 {
		qb.Where(conds[0])
	}
	for i := 1; i < len(conds); i++ {
		qb.Or(conds[i])
	}

	sql := qb.String()

	// 执行SQL语句
	fmt.Println(sql)
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&articles)

	RemoveDuplicates(&articles)

	return
}

func RemoveDuplicates(articles *[]*Article) {
	foundId := make(map[int64]bool)
	j := 0
	for i, x := range *articles {
		if !foundId[x.Id] {
			foundId[x.Id] = true
			(*articles)[j] = (*articles)[i]
			j++
		}
	}
	*articles = (*articles)[:j]
}

func SetupTotalArticleCount() (err error) {
	o := orm.NewOrm()
	totalArticleCount, err = o.QueryTable("article").Count()
	if err != nil {
		return
	}
	beego.Info("Total article count = ", totalArticleCount)
	return
}

func GetTotalArticleCount() (int64) {
	return totalArticleCount
}
