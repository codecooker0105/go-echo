package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/triaton/forum-backend-echo/common/models"
)

type Comment struct {
	models.Base
	Content string
	UserID  uuid.UUID
	BlogID  uuid.UUID
}
