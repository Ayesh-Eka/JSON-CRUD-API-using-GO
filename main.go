package main

import (
	"github.com/ayesh2/go-crud2/controllers"
	"github.com/ayesh2/go-crud2/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadCnvVariables()
	initializers.ConnectToDB()
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