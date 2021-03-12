package blog

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/triaton/forum-backend-echo/common"
	"net/http"
)

type (
	Controller struct {
	}

	AddBlogRequest struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}
)

func (controller Controller) Routes() []common.Route {
	return []common.Route{
		{
			Method:     echo.POST,
			Path:       "/blogs",
			Handler:    controller.AddBlog,
			Middleware: []echo.MiddlewareFunc{common.JwtMiddleWare()},
		},
	}
}

func (controller Controller) AddBlog(ctx echo.Context) error {
	params := new(AddBlogRequest)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*common.JwtCustomClaims)
	//db := database.GetInstance()
	//var blog models.Blog
	//blog.Content = params.Content
	//blog.Title = params.Title
	//blog.UserID = claims.Id
	//db.Create(&blog)
	return ctx.JSON(http.StatusOK, claims)
}
