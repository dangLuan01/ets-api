package auth

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateAccessToken(user models.User) (string, error)
	GenerateRefreshToken(user models.User) (RefreshToken, error)
	ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error)
	DecryptAccessTokenPayload(tokenString string) (*v1dto.EncryptedPayload, error)
	StoreRefreshToken(token RefreshToken) error
	ValidaRefreshToken(token string) (RefreshToken, error)
	RevokeRefreshToken(token string) error
}