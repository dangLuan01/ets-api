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
	UpdateUser(ctx *gin.Context, params v1dto.UpdateUserInput) error
	DeleteUser(uuid uuid.UUID) error
	UpdatePassword(ctx *gin.Context, params v1dto.UpdatePasswordInput) error
	GetAttemptByUserUUID(ctx *gin.Context, params v1dto.GetAttemptByUserUuuidParams) ([]v1dto.UserAttemptDTO, int64, error)
}