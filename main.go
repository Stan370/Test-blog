package main

import (
	"net/http"

	"github.com/Stan370/Test-blog/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//创建gorm.DB对象的时候连接并没有被创建，在具体使用的时候才会创建。gorm内部，准确的说是database/sql内部会维护一个连接池，可以通过参数设置最大空闲连接数，连接最大空闲时间等。使用者不需要管连接的创建和关闭。
	dbObject := db.InitDatabase()
	postCtrl := &api.postController{
		Dbconn: dbObject,
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})
	v1 := r.Group("/api/v1", gin.Logger())
	{
		v1.GET("/posts", postCtrl.getAllPosts)
		v1.GET("/posts/:id", postCtrl.getPostByID)
		v1.POST("/posts", postCtrl.createPost)
		v1.PUT("/posts/:id", postCtrl.updatePost)
		v1.DELETE("/posts/:id", postCtrl.deletePost)
		//v1.GET("/healthcheck", healthCheckHandler)
	}

	r.Run(":8080")
}
