package blogs

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/triaton/forum-backend-echo/blogs/models"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/database"
	"net/http"
)

type (
	CommentController struct {
	}

	AddCommentRequest struct {
		Content string    `json:"content" validate:"required"`
		BlogId  uuid.UUID `json:"blogId" validate:"required"`
	}
)

func (controller CommentController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.GET,
			Path:    "/blogs/:blogId/comments",
			Handler: controller.GetComments,
		},
		{
			Method:     echo.POST,
			Path:       "/blogs/:blogId/comments",
			Handler:    controller.AddComment,
			Middleware: []echo.MiddlewareFunc{common.JwtMiddleWare()},
		},
	}
}

func (controller CommentController) GetComments(ctx echo.Context) error {
	blogId := ctx.Param("blogId")
	db := database.GetInstance()
	var comments []models.Comment
	db.Find(&comments, "blog_id = ?", blogId)
	return ctx.JSON(http.StatusOK, comments)
}

func (controller CommentController) AddComment(ctx echo.Context) error {
	params := new(AddCommentRequest)
	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*common.JwtCustomClaims)
	db := database.GetInstance()
	var blog models.Blog
	err := db.First(&blog, "id = ?", params.BlogId).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Blog not found")
	}
	comment := models.Comment{
		Content: params.Content,
		UserID:  claims.Id,
		BlogID:  params.BlogId,
	}
	db.Create(&comment)
	return ctx.JSON(http.StatusOK, comment)
}
