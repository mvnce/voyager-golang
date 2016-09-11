package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
	"math/rand"
	//"crypto/rand"
	//"crypto/sha1"
	//"io"
	//"fmt"
	//"encoding/base64"
	//"bytes"
	//"hash"
)

const (
	SaltSize = 16
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`~!@#$%^&*()-=_+"
	RandMax = 77
)

func randomSalt() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, SaltSize)
	for i := range b {
		b[i] = chars[rand.Intn(RandMax)]
	}
	return string(b)
}

type User struct {
	Id 	int64	`json:"id"`
	Name 	string	`json:"name"`
	Password string	`json:"password"`
	Salt 	string	`json:"salt"`
	Email 	string	`json:"email"`
	Type 	string	`json:"type"`
	Joined 	time.Time `json:"joined"`
}

func SignUp() error {
	//user.Joined = time.Now()
	//
	//o := orm.NewOrm()
	//
	//salt := randomSalt()
	//
	//sha
	//
	//hash.Hash64(salt)
	//
	//user.Salt = salt
	//
	//_, err := o.Insert(&p)
	//
	//if err == nil {
	//	return err
	//}

	return nil
}

func SignIn() error {
	return nil
}

func SignOut() error {
	return nil
}
