package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id 	int64
	UserId 	int64
	Title 	string
	Content string
	Status 	string
	Created time.Time
	Updated time.Time
}

func AddPost(post Post) error {
	o := orm.NewOrm()

	_, err := o.Insert(&post)

	if err == nil {
		return err
	}

	return nil
}
