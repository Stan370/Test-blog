package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID       int `gorm:"unique;not null"`
	Username string
	Email    string
	Password string
	Posts    []Post `gorm:"ForeignKey:PostID"`
}

type Post struct {
	PostID    uint   `gorm:"unique;not null"`
	Title     string `gorm:"type:text;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   //Foreign key
	CreatedAt time.Time
}

func CreatePost(db *gorm.DB, title, content string, authorID uint) error {
	if title == "" || content == "" {
		return errors.New("title and content cannot be empty")
	}

	post := Post{
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
	}

	// Attempt to create the post
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&post)
	if result.Error != nil {
		return fmt.Errorf("failed to create post: %w", result.Error)
	}

	// Check the number of rows affected after creating the post
	if result.RowsAffected == 0 {
		return errors.New("post already exists or no changes were made")
	}

	return result.Error
}

func InitDatabase() *gorm.DB {
	var users = User{ID: 1000, Username: "Jiaming", Email: "kjmcs2048@gmail.com"}

	db, err := gorm.Open(mysql.Open("blog.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln("Fail to connect database, please check config")
		return nil
	}

	err = CreatePost(db, "New Title", "This is a new post content.", 1000)
	if err != nil {
		log.Panicln("Failed to create post: " + err.Error())
	}

	// AutoMigrate will create the tables based on the models if they don't exist
	if err := db.AutoMigrate(&users, &Post{}); err != nil {
		log.Panicln("Error auto migrating models: " + err.Error())
		return nil
	}

	return db
}
