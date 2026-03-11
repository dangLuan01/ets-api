package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx *gin.Context, email, password string) (string, string, int, error)
	Logout(ctx *gin.Context, refreshTokenString string) error
	RefreshToken(ctx *gin.Context, token string) (string, string, int, error)
	Register(ctx *gin.Context, input v1dto.RegisterInput) error
	RegisterOTP(ctx *gin.Context, code string) error
}