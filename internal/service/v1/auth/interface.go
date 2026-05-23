package v1service

import (
	"context"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx *gin.Context, params v1dto.LoginInput) (string, string, int, error)
	Logout(ctx *gin.Context, refreshTokenString string) error
	RefreshToken(ctx *gin.Context, token string) (string, string, int, error)
	Register(ctx *gin.Context, input v1dto.RegisterInput) error
	Oauth2Login(ctx context.Context, provider string) (string, error)
	Oauth2CallBack(ctx context.Context, provider, code, state, error string) (string, string, int, error)
}