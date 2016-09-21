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

func AddUser(user User) (int64, error) {
	user.Joined = time.Now()

	o := orm.NewOrm()

	return o.Insert(&user)
}

func VerifyCredential(user User) int64 {
	dbUser := new(User)

	o := orm.NewOrm()
	qs := o.QueryTable("voyager_user")

	err := qs.Filter("email", user.Email).One(dbUser)

	if err != nil {
		return 0
	}

	if user.Password == dbUser.Password {
		return dbUser.Id
	}

	return 0
}
