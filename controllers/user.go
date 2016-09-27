package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"strings"
)

const (
	hmacKey = "Vincent,HmacKey,Sample"
)

type CustomClaims struct {
	UserId int64
	Email string
	jwt.StandardClaims
}

type UserController struct {}

func (uc UserController) SignUp(context *gin.Context) {
	var user models.User
	context.Bind(&user)
	userId, err := models.AddUser(user)

	if err == nil {

		claims := CustomClaims{
			userId,
			user.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
				Issuer:    "test",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(hmacKey))

		fmt.Printf("%v %v", tokenStr, err)

		if err != nil {
			return
		}

		context.JSON(200, gin.H{"message": "ok", "data": map[string]string { "token": tokenStr } })
	} else {
		context.JSON(401, gin.H{"message": "authentication failed", "data": map[string]string {} })
	}
}

func (uc UserController) SignIn(context *gin.Context) {
	var user models.User
	context.Bind(&user);
	userId := models.VerifyCredential(user)

	if userId > 0 {
		claims := CustomClaims{
			userId,
			user.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
				Issuer:    "test",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(hmacKey))

		if err != nil {
			return
		}

		context.JSON(200, gin.H{"message": "ok", "data": map[string]string { "token": tokenStr } })
	} else {
		context.JSON(401, gin.H{"message": "authentication failed", "data": map[string]string {} })
	}
}
func (uc UserController) Validate(context *gin.Context) {
	var authHeader = context.Request.Header.Get("Authorization")
	var tokens = strings.Split(authHeader, " ")

	token, err := jwt.ParseWithClaims(tokens[1], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Println(claims.UserId);

		context.JSON(200, gin.H{"message": "ok", "data": map[string]string { "validate": "ok" } })
	} else {
		fmt.Println(err)
		context.JSON(401, gin.H{"message": "authentication failed", "data": map[string]string {} })
	}
}

// check freshness of current token
func CheckToken(tk string) bool {
	token, err := jwt.ParseWithClaims(tk, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Println(claims.UserId);
		return true;
	} else {
		fmt.Println(err)
		return false;
	}
}

func GetUserId(tk string) int64 {
	token, err := jwt.ParseWithClaims(tk, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserId;
	} else {
		fmt.Println(err)
		return 0;
	}
}
