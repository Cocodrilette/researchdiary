package models

import (
	"testing"

	"github.com/cocodrilette/researchdiary/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

/*
*
* Infra setup
*
 */

type TestDB struct {
	conn *gorm.DB
}

func (db *TestDB) Connect() (*gorm.DB, error) {
	conn, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err == nil {
		db.conn = conn
	}

	return conn, err
}

var TestDatabase = TestDB{}

/*
*
* Tests
*
 */

func TestArticleManager_CRUD(t *testing.T) {
	db, err := TestDatabase.Connect()
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	err = db.AutoMigrate(&Article{}, &models.Author{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	manager := &ArticleManager{DB: db}

	// Clean up
	db.Exec("DELETE FROM articles")
	db.Exec("DELETE FROM authors")

	var article Article

	t.Run("Create", func(t *testing.T) {
		article = Article{Title: "Test Title", Author: models.Author{FirstName: "John", LastName: "Doe"}}
		err := manager.Create(&article)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}
		if article.ID == 0 {
			t.Error("Article ID not set after create")
		}
	})

	t.Run("Find", func(t *testing.T) {
		articles, err := manager.Find(nil)
		if err != nil {
			t.Fatalf("Find failed: %v", err)
		}
		if len(articles) != 1 {
			t.Errorf("Expected 1 article, got %d", len(articles))
		}
	})

	t.Run("Update", func(t *testing.T) {
		title := "Updated Title"
		article.Title = title

		manager.Save(&article)

		// Verify update
		var updated Article
		res := manager.DB.First(&updated, article.ID)
		if res.Error != nil || updated.Title != "Updated Title" {
			t.Error("Update not reflected")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		// Delete the article
		result := manager.Delete(&article)
		if result.Error != nil {
			t.Fatalf("Delete failed: %v", result.Error)
		}

		// Verify delete
		var deleted Article
		res := manager.DB.First(&deleted, article.ID)
		if res.Error == nil {
			t.Error("Delete not successful: article still found")
		} else if res.Error != gorm.ErrRecordNotFound {
			t.Errorf("Unexpected error when verifying delete: %v", res.Error)
		}
	})
}
