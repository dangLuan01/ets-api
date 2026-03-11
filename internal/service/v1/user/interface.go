package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	GetUserByUUID(ctx *gin.Context) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(uuid uuid.UUID, user models.User) (models.User, error)
	DeleteUser(uuid uuid.UUID) error
	CheckStatus(ctx *gin.Context, uuid string) error
	ChangePassword(ctx *gin.Context, params v1dto.ChangerPasswordParams) error
	UpdateCountUpload(uuid string) error
}