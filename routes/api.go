package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	Auth "github.com/triaton/forum-backend-echo/controllers/api/auth"
	Blogs "github.com/triaton/forum-backend-echo/controllers/api/blog"
	Users "github.com/triaton/forum-backend-echo/controllers/api/users"
)

func DefineApiRoute(e *echo.Echo) {
	controllers := []common.Controller{
		Auth.Controller{},
		Users.Controller{},
		Blogs.Controller{},
	}
	var routes []common.Route
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}
	api := e.Group("/api/v1")
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
