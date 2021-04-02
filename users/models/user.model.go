package models

import (
	BlogModels "github.com/triaton/go-echo-boilerplate/blogs/models"
	"github.com/triaton/go-echo-boilerplate/common"
	CommonModels "github.com/triaton/go-echo-boilerplate/common/models"
	"github.com/triaton/go-echo-boilerplate/common/utils"
)

type User struct {
	CommonModels.Base
	Email    string `gorm:"type:varchar(100);unique_index"`
	Name     string
	Role     common.UserRole
	Password string
	Blogs    []BlogModels.Blog
}

func (user User) String() string {
	return user.Name
}

func (user *User) BeforeSave() (err error) {
	hashed, err := utils.GetPasswordUtil().HashPassword(user.Password)
	user.Password = hashed
	return
}
