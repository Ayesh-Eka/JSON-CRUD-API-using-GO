package controllers

import (
	"github.com/ayesh2/go-crud2/initializers"
	"github.com/ayesh2/go-crud2/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
    //Get data off request body
    var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//Create a Post
    post := models.Post{Title: body.Title, Body: body.Body}

    result := initializers.DB.Create(&post)

    if result.Error !=nil{
		c.Status(400)
		return
	}


	//return it


	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
    initializers.DB.Find(&posts)


	//Respond
    c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	//Get id off url
    id := c.Param("id")

	//Get the posts
	var post models.Post
    initializers.DB.First(&post, id)


	//Respond
    c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//Get the id of url
    id := c.Param("id")

	//Get the data off request body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
    var post models.Post
    initializers.DB.First(&post, id)

	//Update it
    initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, })

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})


}

func PostsDelete(c *gin.Context) {
	//Get the id
    id := c.Param("id")

	//Delete
	initializers.DB.Delete(&models.Post{}, id)


	//Respond
    c.Status(200)

}