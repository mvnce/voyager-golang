package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"github.com/gin-gonic/contrib/sessions"
	//"strconv"
	//"voyager-golang/forms"
	"fmt"
)

type UserController struct {}

//func getId(context *gin.Context) (id int64, err error) {
//	session := sessions.Default(context)
//	sid := session.Get("user_id")
//
//	if sid != nil {
//		id, err = strconv.ParseInt(sid, 10, 64)
//
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	return id, err
//}

func (uc UserController) SignUp(context *gin.Context) {
	fmt.Println("Signup Controller")
	var user models.User

	if context.Bind(&user) != nil {
		context.JSON(406, gin.H{"status": "error"})
		return
	}

	ret := models.SignUp(user)

	if ret == nil {
		session := sessions.Default(context)
		session.Set("user_id", user.Id)
		session.Set("user_name", user.Name)
		session.Save()
		context.JSON(200, gin.H{"message": "ok", "data": user})
	} else {
		context.JSON(406, gin.H{"message": "error", "error": "bad signup"})
	}
}

func (uc UserController) SignIn(context *gin.Context) {
	models.SignIn()
}

func (uc UserController) SignOut(context *gin.Context) {
	models.SignOut()
}
