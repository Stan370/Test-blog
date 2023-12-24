package api

import (
	"net/http"

	"github.com/Stan370/Test-blog/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postController struct {
	Dbconn *gorm.DB
}

func (controller *postController) getAllPosts(c *gin.Context) {
	var Post db.Post
	if err := controller.Dbconn.Find(&Post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Post)
	return
}

func (controller *postController) getPostByID(c *gin.Context) {
	id := c.Param("id")
	var post db.Post
	if err := controller.Dbconn.Where("post_id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
	return
}

func (controller *postController) createPost(c *gin.Context) {
	var newPost db.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate unique ID for new post
	// Append newPost to 'posts'
	// Return the created post or appropriate message
}

func updatePost(c *gin.Context) {
	//To be add
}

func deletePost(c *gin.Context) {
	//To be add
}
