package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id int64
	UserId int64	`json:"user_id"`
	PostId int64	`json:"post_id"`
	Content string	`json:"content"`
	Status string	`json:"status"`
	Created time.Time
	Updated time.Time
}

func AddComment(comment Comment) error {
	comment.Created = time.Now()
	comment.Updated = time.Now()

	o := orm.NewOrm()

	_, err := o.Insert(&comment)

	return err
}

// Get all comments for a particular thread
// para: post_id
func GetComments(pid int64) ([]*Comment, error) {
	o := orm.NewOrm()

	comments := make([]*Comment, 0)

	qs := o.QueryTable("voyager_comment")
	_, err := qs.OrderBy("updated").Filter("post_id", pid).All(&comments)

	return comments, err
}