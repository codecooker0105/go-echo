package main

import (
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/triaton/go-echo-boilerplate/common"
	"github.com/triaton/go-echo-boilerplate/database"
	"github.com/triaton/go-echo-boilerplate/routes"
	"log"
)

func main() {
	// Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	m := database.GetMigrations(db)
	err = m.Migrate()
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
	server.Logger.Fatal(server.Start(":1200"))
}
