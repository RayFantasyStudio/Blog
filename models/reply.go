package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Id           int64
	ArticleId    int64
	UserName     string
	Content      string`orm:"size(500)"`
	Time         time.Time
	LikeCount    int64
	QuoteReplyId int64
}
func AddReply(reply *Reply) error {
	reply.Time = time.Now()
	o := orm.NewOrm()
	_, err := o.Insert(reply)
	return err
}
func DeleteReply(Id int64) error {
	o := orm.NewOrm()
	reply := new(Reply)
	var err error
	if Id > -1 {
		reply.Id = Id
		_, err = o.Delete(reply)
	}
	return err
}
func GetReplies(aid int64) ([]*Reply,error){
	o := orm.NewOrm()
	qs := o.QueryTable("reply").Filter("article_id",aid)
	replies := make([]*Reply,0)
	_,err :=  qs.All(&replies)
	return replies,err
}
func GetAdminReplies() ([]*Reply,error){
	o := orm.NewOrm()
	qs := o.QueryTable("reply")
	replies := make([]*Reply,0)
	_,err :=  qs.All(&replies)
	return replies,err
}