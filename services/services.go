package services

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"strconv"
)

func AddPost(context *gin.Context) {
	var post models.Post
	context.BindJSON(&post)

	ret := models.AddPost(post)

	if ret == nil {
		context.JSON(201, gin.H{"status": "ok" ,"data": post})
	} else {
		context.JSON(400, gin.H{"error": "bad post input"})
	}
}


func GetPosts(context *gin.Context) {
	posts, err := models.GetPosts()

	if err == nil {
		context.JSON(200, gin.H{"status": "ok" ,"data": posts})
	} else {
		context.JSON(404, gin.H{"error": "no post(s) in the table"})
	}
}

func GetPost(context *gin.Context) {
	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	post, err := models.GetPost(id)

	if err == nil {
		context.JSON(200, gin.H{"status": "ok" ,"data": post})
	} else {
		context.JSON(404, gin.H{"error": "no post in the table"})
	}
}

func UpdatePost(context *gin.Context) {
	var post models.Post
	context.Bind(&post)

	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	ret := models.UpdatePost(id, post)

	if ret == nil {
		context.JSON(201, gin.H{"status": "ok" ,"data": post})
	} else {
		context.JSON(400, gin.H{"error": "bad post input"})
	}
}

func DeletePost(context *gin.Context) {
	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	ret := models.DeletePost(id)

	if ret == nil {
		context.JSON(200, gin.H{"status": "id # " + sid + " has been deleted"})
	} else {
		context.JSON(404, gin.H{"error": "user not found"})
	}
}
