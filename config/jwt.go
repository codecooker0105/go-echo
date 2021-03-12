package config

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/triaton/forum-backend-echo/common"
)

const TokenExpiresIn = 3 // specify token expire hours

func JwtConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &common.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
}
