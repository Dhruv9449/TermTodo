package database

import (
	"log"

	"github.com/Dhruv9449/TermTodo/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global database object
var DB *gorm.DB

// Function to connect to the database
func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("termtodo.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Migrate the database
	DB.AutoMigrate(
		&models.TodoItem{},
		&models.TodoTable{},
	)
}
