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

type AuthService interface {
	Login(ctx *gin.Context, email, password string) (string, string, int, error)
	Logout(ctx *gin.Context, refreshTokenString string) error
	RefreshToken(ctx *gin.Context, token string) (string, string, int, error)
	Register(ctx *gin.Context, input v1dto.RegisterInput) error
	RegisterOTP(ctx *gin.Context, code string) error
}

type ExamService interface {
	FindExamById(id string) (models.Exam, error)
}