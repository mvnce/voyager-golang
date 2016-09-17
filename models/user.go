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

type Token struct {
	Token string `json:"token"`
}

func AddUser(user User) error {
	user.Joined = time.Now()

	o := orm.NewOrm()

	_, err := o.Insert(&user)

	return err
}

func VerifyCredential(user User) bool {
	return true
}
