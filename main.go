package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"voyager-golang/models"
	"github.com/astaxie/beego/orm"
	"voyager-golang/controllers"
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
		post := new(controllers.PostController)

		v1.GET("/posts", post.GetPosts)
		v1.GET("/posts/:id", post.GetPost)
		v1.POST("/posts", post.AddPost)
		v1.PUT("/posts/:id", post.UpdatePost)
		v1.DELETE("/posts/:id", post.DeletePost)

		// curl -i http://localhost:8080/api/v1/posts
		// http -f GET http://localhost:8080/api/v1/posts "Authorization:Bearer xxxxxxxxx"  "Content-Type: application/json"

		// curl -i http://localhost:8080/api/v1/posts/1
		// curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_id\": 5, \"title\": \"First Title\", \"content\": \"Content Field\", \"status\": \"posted\"}" http://localhost:8080/api/v1/posts
		// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"status\": \"updated\" }" http://localhost:8080/api/v1/posts/1
		// curl -i -X DELETE http://localhost:8080/api/v1/posts/6

		user := new(controllers.UserController)

		v1.POST("/auth/signup", user.SignUp)
		v1.POST("/auth/signin", user.SignIn)
		v1.POST("/auth/check", user.Check)

		// curl -i -X POST -H "Content-Type: application/json" -d "{\"email\": \"facebowl\", \"password\": \"facebowl\"}" http://localhost:8080/api/v1/auth/signup
		// curl -i -X POST -H "Content-Type: application/json" -d "{\"token\": \"token_sample\"}" http://localhost:8080/api/v1/user/checkstatus


		comment := new(controllers.CommentController)

		v1.POST("/comments", comment.AddComment)
		v1.GET("/comments/:id", comment.GetComments)

		// curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_id\": 1, \"post_id\": 19, \"content\": \"Content Field\", \"status\": \"new\"}" http://localhost:8080/api/v1/comments
		// curl -i http://localhost:8080/api/v1/comments/19

	}

	router.Run(":8080")
}
