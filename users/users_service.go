package users

import (
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/users/models"
)

func FindUserByEmail(email string) *models.User {
	db := database.GetInstance()
	var user models.User
	err := db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}
	return nil
}
