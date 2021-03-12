package models

import (
	_ "github.com/jinzhu/gorm"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/common/utils"
)

type User struct {
	Base
	Email    string `gorm:"type:varchar(100);unique_index"`
	Name     string
	Role     common.UserRole
	Password string
	Blogs    []Blog
	Comments []Comment
}

func (user User) String() string {
	return user.Name
}

func (user *User) BeforeSave() (err error) {
	hashed, err := utils.HashPassword(user.Password)
	user.Password = hashed
	return
}
