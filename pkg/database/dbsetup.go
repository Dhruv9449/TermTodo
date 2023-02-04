package database

import (
	"log"

	"github.com/Dhruv9449/TermTodo/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("termtodo.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	DB.AutoMigrate(
		&models.TodoItem{},
		&models.TodoTable{},
	)
}
