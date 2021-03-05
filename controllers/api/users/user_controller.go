package users

import (
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	"net/http"
)

type Controller struct{}

func (controller Controller) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.GET,
			Path:    "/users/bll",
			Handler: controller.Index,
		},
		{
			Method:  echo.GET,
			Path:    "/users/all",
			Handler: controller.All,
		},
	}
}

func (controller Controller) All(c echo.Context) error {
	return c.String(http.StatusOK, "All users from api")
}

func (controller Controller) Index(c echo.Context) error {
	return c.String(http.StatusOK, "index from api")
}
