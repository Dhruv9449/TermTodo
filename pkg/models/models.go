package models

import "time"

type TodoTable struct {
	ID    uint `gorm:"primary_key"`
	Title string
	Items []TodoItem `gorm:"ForeignKey:TodoTableID"`
}

type TodoItem struct {
	ID          uint      `gorm:"primary_key"`
	CreatedOn   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Title       string
	Description string
	Completed   bool `gorm:"default:false"`
	TodoTableID uint
}
