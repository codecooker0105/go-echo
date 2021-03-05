package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	"net/http"
	"os"
	"time"
)

type Controller struct {
}

func (controller Controller) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/auth/login",
			Handler: controller.Login,
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

// Get login information as well as user JWT token.
func (controller Controller) Login(c echo.Context) error {
	var (
		email, password string
	)

	email = c.FormValue("email")
	password = c.FormValue("password")

	// Make a checking in database instead
	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusForbidden, "Please provide email and password credentials")
	}
	print("result=%s:%s", email, password)
	if email != "honglin328@yandex.com" || password != "secret" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authentication failed")
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "honglin"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
