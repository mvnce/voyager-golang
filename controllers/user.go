package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strconv"
	"fmt"
)

const (
	hmacKey = "Vincent,HmacKey,Sample"
)

type UserController struct {}

func (uc UserController) SignUp(context *gin.Context) {
	var user models.User
	context.Bind(&user)

	newId, err := models.AddUser(user)

	if err == nil {
		expTime := time.Now().Add(time.Minute * 5).Unix()

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": user.Name,
				"exp": expTime,

			})

		tokenStr, err := token.SignedString([]byte(hmacKey))

		if err != nil {
			return
		}

		data := map[string]string {
			"token": tokenStr,
			"id": strconv.FormatInt(newId, 10),
		}

		context.JSON(200, gin.H{"message": "ok", "data": data})
	} else {
		context.JSON(406, gin.H{"message": "error", "error": "bad signup"})
	}
}

func (uc UserController) SignIn(context *gin.Context) {
	var user models.User

	context.Bind(&user);

	ret := models.VerifyCredential(user)

	if ret {
		expTime := time.Now().Add(time.Minute * 5).Unix()

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": user.Name,
				"exp": expTime,

			})

		tokenStr, err := token.SignedString([]byte(hmacKey))

		if err != nil {
			return
		}

		data := map[string]string {
			"token": tokenStr,
			"exp": strconv.FormatInt(expTime, 10),
		}

		context.JSON(200, gin.H{"message": "ok", "data": data})
	} else {

		data := map[string]string {}
		context.JSON(200, gin.H{"message": "authentication failed", "data": data})
	}
}

func (uc UserController) Check(context *gin.Context) {
	var tokenModel models.Token

	if context.Bind(&tokenModel) != nil {
		context.JSON(406, gin.H{"status": "error"})
		return
	}

	token, _ := jwt.Parse(tokenModel.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["userid"], claims["exp"])

		context.JSON(200, gin.H{"message": "ok", "data": "validated"})
	} else {
		context.JSON(406, gin.H{"message": "error", "error": "check status failed"})
	}
}
