package controllers

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"strconv"
)

type CommentController struct {}

func (cc CommentController) AddComment(context *gin.Context) {
	var comment models.Comment
	context.Bind(&comment)

	ret := models.AddComment(comment)

	if ret == nil {
		context.JSON(201, gin.H{"message": "ok" ,"data": comment})
	} else {
		context.JSON(400, gin.H{"error": "bad post input"})
	}
}

func (cc CommentController) GetComments(context *gin.Context) {
	sid := context.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	comments, err := models.GetComments(id)

	if err == nil {
		context.JSON(200, gin.H{"message": "ok" ,"data": comments})
	} else {
		context.JSON(404, gin.H{"error": "no comments(s) in the table"})
	}
}