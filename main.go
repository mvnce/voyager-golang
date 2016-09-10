package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"voyager-golang/models"
	"github.com/astaxie/beego/orm"
	"voyager-golang/services"
	"fmt"
)

func init()  {
	models.InitDB()
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}

func Cors() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	//	c.Next()
	//}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	router := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("session", store))

	router.Use(Cors())

	v1 := router.Group("api/v1")
	{
		v1.GET("/posts", services.GetPosts)
		// 104.131.139.229:8080
		// curl -i http://104.131.139.229:8080/api/v1/posts

		v1.GET("/posts/:id", services.GetPost)
		// curl -i http://localhost:8080/api/v1/posts/1

		v1.POST("/posts", services.AddPost)
		// curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_id\": 5, \"title\": \"First Title\", \"content\": \"Content Field\", \"status\": \"posted\"}" http://localhost:8080/api/v1/posts

		v1.PUT("/posts/:id", services.UpdatePost)
		// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"status\": \"updated\" }" http://localhost:8080/api/v1/posts/1

		v1.DELETE("/posts/:id", services.DeletePost)
		// curl -i -X DELETE http://localhost:8080/api/v1/posts/6
	}

	router.Run(":8080")
}
