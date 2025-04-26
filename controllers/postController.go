package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shamsher4499/go-crud/initializers"
	"github.com/shamsher4499/go-crud/models"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create post",
		})
	}

	c.JSON(200, gin.H{
		"message":  "Post created",
		"response": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message":  "Posts fetched",
		"response": posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"message":  "Post fetched",
		"response": post,
	})
}

func UpdatePost(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Post not found",
		})
	}

	response := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if response.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to update post",
		})
	}

	c.JSON(200, gin.H{
		"message":  "Post updated",
		"response": post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Post not found",
		})
	}

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
