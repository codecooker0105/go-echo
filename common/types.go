package common

import (
	"github.com/labstack/echo/v4"
	_ "net/http"
)

type Route struct {
	Method     string
	Path       string
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

type Controller interface {
	Routes() []Route
}
