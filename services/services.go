package services

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"strconv"
)

func AddPost(c *gin.Context) {
	var post models.Post
	c.Bind(&post)

	ret := models.AddPost(post)

	if ret == nil {
		c.JSON(201, gin.H{"status": "ok" ,"data": post})
	} else {
		c.JSON(400, gin.H{"error": "bad post input"})
	}
}


func GetPosts(c *gin.Context) {
	//var posts []models.User

	posts, err := models.GetPosts()

	if err == nil {
		c.JSON(200, gin.H{"status": "ok" ,"data": posts})
	} else {
		c.JSON(404, gin.H{"error": "no post(s) in the table"})
	}
}

func GetPost(c *gin.Context) {
	sid := c.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	post, err := models.GetPost(id)

	if err == nil {
		c.JSON(200, gin.H{"status": "ok" ,"data": post})
	} else {
		c.JSON(404, gin.H{"error": "no post in the table"})
	}
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	c.Bind(&post)

	sid := c.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	ret := models.UpdatePost(id, post)

	if ret == nil {
		c.JSON(201, gin.H{"status": "ok" ,"data": post})
	} else {
		c.JSON(400, gin.H{"error": "bad post input"})
	}
}

func DeletePost(c *gin.Context) {
	sid := c.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	ret := models.DeletePost(id)

	if ret == nil {
		c.JSON(200, gin.H{"status": "id # " + sid + " has been deleted"})
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
