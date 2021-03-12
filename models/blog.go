package models

import (
	_ "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Blog struct {
	Base
	Title   string
	Content string
	UserID  uuid.UUID
}
