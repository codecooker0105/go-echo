package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/common/utils"
	"github.com/triaton/forum-backend-echo/users"
	"net/http"
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
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if user := users.GetUsersService().FindUserByEmail(params.Email); user != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "email is already used")
	}
	user := users.GetUsersService().AddUser(params.Name, params.Email, params.Password)
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

	user := users.GetUsersService().FindUserByEmail(params.Email)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	if matched := utils.CheckPasswordHash(params.Password, user.Password); !matched {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	// Create token
	token, err := AuthService().GetAccessToken(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
