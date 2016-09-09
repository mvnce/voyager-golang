package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"voyager-golang/models"
	"github.com/astaxie/beego/orm"
	"voyager-golang/services"
)

func init()  {
	models.InitDB()
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	router := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("session", store))

	router.Use(Cors())

	v1 := router.Group("api/v1")
	{
		//v1.GET("/posts", Gettings)
		//v1.GET("/posts/:id", Getting)
		
		v1.POST("/posts", services.AddPost)
		// curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_id\": 5, \"title\": \"First Title\", \"content\": \"Content Field\", \"status\": \"posted\"}" http://localhost:8080/api/v1/posts

		//v1.PUT("/posts/:id", Updating)
		//v1.DELETE("/posts/:id", Deleting)
	}


	router.Run(":8080")

	router.Run(":8080")
}
