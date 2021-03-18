package server

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/migrations"
	"github.com/triaton/forum-backend-echo/routes"
)

func MakeServer() *echo.Echo {
	//Define API wrapper
	api := echo.New()
	api.Validator = &common.CustomValidator{Validator: validator.New()}
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	// CORS middleware for API endpoint.
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	db := database.GetInstance()
	m := migrations.GetMigrations(db)
	err := m.Migrate()
	if err == nil {
		print("Migrations did run successfully")
	} else {
		print("migrations failed.", err)
	}
	routes.DefineApiRoute(api)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		if req.URL.Path[:4] == "/api" {
			api.ServeHTTP(res, req)
		}

		return
	})
	return server
}
