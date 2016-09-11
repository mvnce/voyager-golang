package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id 	int64 `json:"id"`
	UserId 	int64	`json:"user_id" binding:"required"`
	Title 	string	`json:"title"`
	Content string	`json:"content"`
	Status 	string	`json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func AddPost(p Post) error {
	p.Created = time.Now()
	p.Updated = time.Now()

	o := orm.NewOrm()

	_, err := o.Insert(&p)

	if err == nil {
		return err
	}

	return nil
}

func GetPosts() ([]*Post, error)  {
	o := orm.NewOrm()

	posts := make([]*Post, 0)

	qs := o.QueryTable("voyager_post")
	_, err := qs.OrderBy("-updated").All(&posts)

	return posts, err
}

func GetPost(id int64) (*Post, error) {
	o := orm.NewOrm()

	post := new(Post)

	qs := o.QueryTable("voyager_post")
	err := qs.Filter("id", id).One(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func UpdatePost(id int64, p Post) error {
	o := orm.NewOrm()

	post := &Post{Id: id}

	if o.Read(post) == nil {
		post.Title = p.Title
		post.Content = p.Content
		post.Status = p.Status
		post.Updated = time.Now()
		o.Update(post, "Title", "Content", "Status", "Updated")
	}

	return nil
}

func DeletePost(id int64) error {
	o := orm.NewOrm()

	post := &Post{Id: id}

	_, err := o.Delete(post)

	return err
}
