package main

import (
	"github.com/Scyther2/go-crud/controllers"
	"github.com/Scyther2/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionToDB()
}
func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
