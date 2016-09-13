package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"strconv"
)

type PostController struct{}

func (pc PostController) AddPost(context *gin.Context) {
	var post models.Post
	context.Bind(&post)

	ret := models.AddPost(post)

	if ret == nil {
		context.JSON(201, gin.H{"message": "ok" ,"data": post})
	} else {
		context.JSON(400, gin.H{"error": "bad post input"})
	}
}


func (pc PostController) GetPosts(context *gin.Context) {
	posts, err := models.GetPosts()

	if err == nil {
		context.JSON(200, gin.H{"message": "ok" ,"data": posts})
	} else {
		context.JSON(404, gin.H{"error": "no post(s) in the table"})
	}
}

func (pc PostController) GetPost(context *gin.Context) {
	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	post, err := models.GetPost(id)

	if err == nil {
		context.JSON(200, gin.H{"message": "ok" ,"data": post})
	} else {
		context.JSON(404, gin.H{"error": "no post in the table"})
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
