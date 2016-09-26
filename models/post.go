package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id 	int64 `json:"id"`
	UserId 	int64	`json:"user_id"`
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

	return err
}

func GetPosts() []orm.Params  {
	o := orm.NewOrm()
	var maps []orm.Params
	var query = `
		SELECT post.id, user.name, post.title, post.content, post.updated, COUNT(comment.id) AS total_comments
		FROM voyager_post AS post
		INNER JOIN voyager_user as user
		ON post.user_id=user.id
		LEFT JOIN voyager_comment as comment
		ON post.id=comment.post_id
		GROUP BY post.id
		ORDER BY post.updated DESC;
	`
	o.Raw(query).Values(&maps)

	return maps
}

func GetPost(pid int64) ([]orm.Params, int64) {
	o := orm.NewOrm()
	var maps []orm.Params
	var query = `
		SELECT * FROM voyager_post AS post
		WHERE post.id=?
	`
	cnt, err := o.Raw(query, pid).Values(&maps)
	if err != nil {
		panic(err)
	}

	return maps, cnt
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
