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
	_, err := qs.All(&posts)

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