package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Stan370/Test-blog/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTest() (*gin.Engine, *gorm.DB) {
	// 创建内存 SQLite 数据库进行测试
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	router := gin.Default()

	postController := api.PostController{Dbconn: db}

	router.GET("/posts", postController.GetAllPosts)
	router.GET("/posts/:id", postController.GetPostByID)
	router.POST("/posts", postController.CreatePost)

	return router, db
}

type Post struct {
	PostID    string `gorm:"unique;not null"`
	Title     string `gorm:"type:text;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   //Foreign key
}

func TestGetAllPosts(t *testing.T) {
	router, db := setupTest()
	defer db.Migrator().DropTable(&Post{})

	db.Create(&Post{Title: "Test Title", Content: "Test Content"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var posts []Post
	json.Unmarshal(w.Body.Bytes(), &posts)

	assert.Equal(t, 1, len(posts))
}

func TestGetPostByID(t *testing.T) {
	router, db := setupTest()
	defer db.Migrator().DropTable(&Post{})

	db.Create(&Post{PostID: "1", Title: "Test Title", Content: "Test Content"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var post Post
	json.Unmarshal(w.Body.Bytes(), &post)

	assert.Equal(t, "Test Title", post.Title)
}

func TestCreatePost(t *testing.T) {
	router, db := setupTest()
	defer db.Migrator().DropTable(&Post{})

	newPost := Post{Title: "New Title", Content: "New Content"}
	payload, _ := json.Marshal(newPost)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(payload)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdPost Post
	json.Unmarshal(w.Body.Bytes(), &createdPost)

	assert.Equal(t, "New Title", createdPost.Title)
}
