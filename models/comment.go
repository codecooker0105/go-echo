package models

import (
	_ "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Comment struct {
	Base
	Content string
	UserID  uuid.UUID
	BlogID  uuid.UUID
}
