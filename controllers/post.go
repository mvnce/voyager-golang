package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"strconv"
	"strings"
)

type PostController struct{}

func (pc PostController) AddPost(context *gin.Context) {
	var post models.Post
	var authHeader = context.Request.Header.Get("Authorization")
	var tokens = strings.Split(authHeader, " ")
	var userId = GetUserId(string(tokens[1]))

	if userId > 0 {
		err := context.BindJSON(&post)
		if err != nil {
			panic(err)
		}
		post.UserId = userId

		ret := models.AddPost(post)

		if ret == nil {
			context.JSON(201, gin.H{"message": "ok", "data": post})
		} else {
			context.JSON(400, gin.H{"error": "bad post input"})
		}
	} else {
		context.JSON(401, gin.H{"message": "authentication failed"})
	}




}


func (pc PostController) GetPosts(context *gin.Context) {
	context.JSON(200, gin.H{" message": "ok" ,"data": models.GetPosts() })
}

func (pc PostController) GetPost(context *gin.Context) {
	sid := context.Params.ByName("id")
	pid, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		panic(err)
	}

	post, cnt := models.GetPost(pid)

	if cnt > 0 {
		var post = post[0]

		context.JSON(200, gin.H{ "message": "ok" ,"data": post })
	} else {
		context.JSON(401, gin.H{" message": "error" })
	}
}

func (pc PostController) UpdatePost(context *gin.Context) {
	var post models.Post
	context.Bind(&post)

	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	ret := models.UpdatePost(id, post)

	if ret == nil {
		context.JSON(201, gin.H{"message": "ok" ,"data": post})
	} else {
		context.JSON(400, gin.H{"error": "bad post input"})
	}
}

func (pc PostController) DeletePost(context *gin.Context) {
	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	ret := models.DeletePost(id)

	if ret == nil {
		context.JSON(200, gin.H{"message": "id # " + sid + " has been deleted"})
	} else {
		context.JSON(404, gin.H{"error": "user not found"})
	}
}
