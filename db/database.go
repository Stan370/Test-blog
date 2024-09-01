package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Stan370/Test-blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User DAO logic not write yet
type User struct {
	ID       uint `gorm:"unique;not null" json:"id,omitempty"`
	Username string
	Email    string
	Password string
}

type Post struct {
	PostID    string `gorm:"primaryKey;not null"`
	Title     string `gorm:"type:text;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `json:"-"` //Foreign key
	User      User   `gorm:"-" json:"postid,omitempty"`
	CreatedAt time.Time
}

func CreatePost(db *gorm.DB, PostID, title, content string, authorID uint) error {
	if title == "" || content == "" {
		return errors.New("title and content cannot be empty")
	}

	post := Post{
		PostID:    PostID,
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

func InitDatabase(config *config.Config) *gorm.DB {
	dsn := config.GetDbConnection()
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalln("Fail to connect database, please check config")
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	//Config connection pool
	sqlDB.SetConnMaxLifetime(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	// var users = User{ID: 1000, Username: "john", Email: "Johndoe@gmail.com", Posts: []Post{post}}
	// AutoMigrate will create the tables based on the models if they don't exist
	if err := db.AutoMigrate(&User{}, &Post{}); err != nil {
		log.Panicln("Error auto migrating models: " + err.Error())
		return nil
	}

	//var post = Post{PostID: uuid.New().String(), Title: "New Title", Content: "This is a new post content.", AuthorID: 1000}
	err = CreatePost(db, "1", "New Title", "This is a new post content.", 1000)
	if err != nil {
		log.Panicln("Failed to create post: " + err.Error())
	}

	return db
}
