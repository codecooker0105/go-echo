package models

import (
	BlogModels "github.com/triaton/forum-backend-echo/blogs/models"
	"github.com/triaton/forum-backend-echo/common"
	CommonModels "github.com/triaton/forum-backend-echo/common/models"
	"github.com/triaton/forum-backend-echo/common/utils"
)

type User struct {
	CommonModels.Base
	Email    string `gorm:"type:varchar(100);unique_index"`
	Name     string
	Role     common.UserRole
	Password string
	Blogs    []BlogModels.Blog
	Comments []BlogModels.Comment
}

func (user User) String() string {
	return user.Name
}

func (user *User) BeforeSave() (err error) {
	hashed, err := utils.GetPasswordUtil().HashPassword(user.Password)
	user.Password = hashed
	return
}
