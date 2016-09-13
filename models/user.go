package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id 	int64	`json:"id"`
	Name 	string	`json:"name"`
	Password string	`json:"password"`
	Salt 	string	`json:"salt"`
	Email 	string	`json:"email"`
	Type 	string	`json:"type"`
	Joined 	time.Time `json:"joined"`
}

func SignUp(user User) error {
	user.Joined = time.Now()

	o := orm.NewOrm()

	_, err := o.Insert(&user)

	return err
}

func SignIn() error {
	return nil
}

func SignOut() error {
	return nil
}
