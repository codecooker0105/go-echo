package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/config"
	UserModels "github.com/triaton/forum-backend-echo/users/models"
	"os"
	"sync"
	"time"
)

type authService struct{}

var singleton *authService
var once sync.Once

func AuthService() *authService {
	once.Do(func() {
		singleton = &authService{}
	})
	return singleton
}

func (s *authService) GetAccessToken(user *UserModels.User) (string, error) {
	claims := &common.JwtCustomClaims{
		Name: user.Name,
		Id:   user.ID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * config.TokenExpiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
