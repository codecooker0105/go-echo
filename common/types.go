package common

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "net/http"
)

type (
	Route struct {
		Method     string
		Path       string
		Handler    echo.HandlerFunc
		Middleware []echo.MiddlewareFunc
	}

	Controller interface {
		Routes() []Route
	}

	CustomValidator struct {
		Validator *validator.Validate
	}

	ValidationError struct {
		namespace string
		field     string
		tag       string
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
