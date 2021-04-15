package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/triaton/go-echo-boilerplate/auth"
	"github.com/triaton/go-echo-boilerplate/blogs"
	"github.com/triaton/go-echo-boilerplate/common"
)

func DefineApiRoute(e *echo.Echo) {
	controllers := []common.Controller{
		auth.AuthController{},
		blogs.BlogsController{},
	}
	var routes []common.Route
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}
	api := e.Group("/api/v0")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}
