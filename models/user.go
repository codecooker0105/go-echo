package models

import (
	_ "github.com/jinzhu/gorm"
	"github.com/triaton/forum-backend-echo/common/utils"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user User) String() string {
	return user.Name
}

func (user *User) BeforeSave() (err error) {
	hashed, err := utils.HashPassword(user.Password)
	user.Password = hashed
	return
}
