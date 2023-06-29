package controllers

import (
	"github.com/Scyther2/go-crud/initializers"
	"github.com/Scyther2/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	// Get id of url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// responde it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
