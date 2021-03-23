package users

import (
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/users/models"
	"sync"
)

type usersService struct{}

var singleton *usersService
var once sync.Once

func UsersService() *usersService {
	once.Do(func() {
		singleton = &usersService{}
	})
	return singleton
}

func (u *usersService) FindUserByEmail(email string) *models.User {
	db := database.GetInstance()
	var user models.User
	err := db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}
	return nil
}

func (u *usersService) AddUser(name string, email string, password string) *models.User {
	user := models.User{
		Name:     name,
		Role:     common.Admin,
		Email:    email,
		Password: password,
	}
	db := database.GetInstance()
	db.Create(&user)
	return &user
}
