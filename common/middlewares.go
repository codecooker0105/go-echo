package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	uuid "github.com/satori/go.uuid"
	"os"
)

type JwtCustomClaims struct {
	Name string    `json:"name"`
	Id   uuid.UUID `json:"id"`
	Role UserRole  `json:"role"`
	jwt.StandardClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWT([]byte(key))
}
