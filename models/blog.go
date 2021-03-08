package models

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

type Blog struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
