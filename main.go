package main

import (
	_ "fmt"
	"github.com/go-playground/validator"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/migrations"
	"github.com/triaton/forum-backend-echo/routes"
	_ "net/http"
)

func main() {
	// Load environment file
	err := godotenv.Load()
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
	m := migrations.GetMigrations(db)
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

	// Start server to listen to port 1200
	server.Logger.Fatal(server.Start(":1200"))
}
