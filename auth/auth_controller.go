package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/common/utils"
	"github.com/triaton/forum-backend-echo/config"
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/users"
	"net/http"
	"os"
	"time"
)

type (
	AuthController struct {
	}

	RegisterUserRequest struct {
		Email    string `json:"email" form:"email" query:"email" validate:"email,required"`
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" form:"email" query:"email" validate:"email,required"`
		Password string `json:"password" validate:"required"`
	}
)

func (r RegisterUserRequest) String() string {
	return fmt.Sprintf("%s, %s, %s", r.Email, r.Name, r.Password)
}

func (controller AuthController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/auth/login",
			Handler: controller.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/auth/register",
			Handler: controller.Register,
		},
		{
			Method:     echo.GET,
			Path:       "/auth/profile",
			Handler:    controller.Profile,
			Middleware: []echo.MiddlewareFunc{common.JwtMiddleWare()},
		},
	}
}

func (controller AuthController) Profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	token := user.Claims.(*common.JwtCustomClaims)
	return c.JSON(http.StatusOK, token)
}

func (controller AuthController) Register(ctx echo.Context) error {
	params := new(RegisterUserRequest)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	db := database.GetInstance()
	user := users.FindUserByEmail(params.Email)
	user.Name = params.Name
	user.Email = params.Email
	user.Role = common.Admin
	user.Password = params.Password
	db.Create(&user)
	return ctx.JSON(http.StatusOK, user)
}

func (controller AuthController) Login(ctx echo.Context) error {
	params := new(LoginRequest)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	user := users.FindUserByEmail(params.Email)
	if user == nil {
		return ctx.String(http.StatusUnauthorized, "Invalid email or password")
	}
	if matched := utils.CheckPasswordHash(params.Password, user.Password); !matched {
		return ctx.String(http.StatusUnauthorized, "Invalid email or password")
	}
	// Create token
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
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
