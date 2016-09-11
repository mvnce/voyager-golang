package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
)

type UserController struct {}

func (uc UserController) SignUp(context *gin.Context) {
	models.SignUp()

}

func (uc UserController) SignIn(context *gin.Context) {
	models.SignIn()
}

func (uc UserController) SignOut(context *gin.Context) {
	models.SignOut()
}
