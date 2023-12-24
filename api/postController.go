package api

import (
	"net/http"

	"github.com/Stan370/Test-blog/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostControllerInterface interface {
	GetAllPosts(c *gin.Context)
	GetPostByID(c *gin.Context)
	GetPostByAuthorId(c *gin.Context)
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type PostController struct {
	Dbconn *gorm.DB
}

func (controller *PostController) GetAllPosts(c *gin.Context) {
	var post db.Post
	if err := controller.Dbconn.Find(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
	return
}

func (controller *PostController) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post db.Post
	if err := controller.Dbconn.Where("post_id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
	return
}

func (controller *PostController) CreatePost(c *gin.Context) {
	var newPost db.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost.PostID = uuid.New().String()

	// Add the new post to the database, AuthorId logic not finished
	if err := db.CreatePost(controller.Dbconn, newPost.PostID, newPost.Title, newPost.Content, 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the post"})
		return
	}

	// Return the created post details
	c.JSON(http.StatusCreated, newPost)
}

func (controller *PostController) GetPostByAuthorId(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, id)
	//To be add
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, id)
	//To be add
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, id)
	//To be add
}
