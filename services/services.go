package services

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"time"
	"fmt"
)

func AddPost(c *gin.Context) {
	var post models.Post
	c.Bind(&post)

	fmt.Println(post)

	post.Created = time.Now()
	post.Updated = time.Now()

	ret := models.AddPost(post)

	if ret == nil {
		c.JSON(201, post)
	} else {
		c.JSON(400, gin.H{"error": "bad post input"})
	}
}

