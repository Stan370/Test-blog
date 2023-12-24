package db

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreatePost(t *testing.T) {
	db, err := gorm.Open(mysql.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	// Migrate the models to create necessary tables
	if err := db.AutoMigrate(&User{}, &Post{}); err != nil {
		t.Fatalf("Error migrating models: %v", err)
	}

	// Test case: Valid post creation
	err = CreatePost(db, "New Title", "This is a new post content.", 1000)
	if err != nil {
		t.Errorf("Failed to create post: %v", err)
	}

	// Test case: Empty title and content
	err = CreatePost(db, "", "", 1000)
	if err == nil {
		t.Error("Expected error for empty title and content, but got nil")
	}

	// Test case: Post already exists
	// Creating the same post again should return an error
	err = CreatePost(db, "New Title", "This is a new post content.", 1000)
	if err == nil {
		t.Error("Expected error for existing post, but got nil")
	}
}
