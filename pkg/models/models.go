package models

import "time"

// Table to store todo items
type TodoTable struct {
	ID    uint `gorm:"primary_key"`
	Title string
	Items []TodoItem `gorm:"ForeignKey:TodoTableID"`
}

// Todo item model
type TodoItem struct {
	ID          uint      `gorm:"primary_key"`
	CreatedOn   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Title       string
	Description string
	Completed   bool `gorm:"default:false"`
	TodoTableID uint
}
