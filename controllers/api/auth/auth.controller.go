package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/common/utils"
	"github.com/triaton/forum-backend-echo/config"
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/models"
	"net/http"
	"os"
	"time"
)

type (
	Controller struct {
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

func (controller Controller) Routes() []common.Route {
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

func (controller Controller) Profile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

func (controller Controller) Register(ctx echo.Context) error {
	params := new(RegisterUserRequest)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	db := database.GetInstance()
	var user models.User
	err := db.First(&user, "email = ?", params.Email).Error
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "registered email")
	}
	user.Name = params.Name
	user.Email = params.Email
	user.Password = params.Password
	db.Create(&user)
	return ctx.JSON(http.StatusOK, user)
}

func (controller Controller) Login(ctx echo.Context) error {
	params := new(LoginRequest)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	db := database.GetInstance()

	user := models.User{}
	err := db.First(&user, "email = ?", params.Email).Error
	if err != nil {
		print(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	if matched := utils.CheckPasswordHash(params.Password, user.Password); !matched {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * config.TokenExpiresIn).Unix()
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
