package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shamsher4499/go-crud/controllers"
	"github.com/shamsher4499/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.Run() // listen and serve on 0.0.0.0:8080
}
