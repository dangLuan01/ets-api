package auth

import (
	"context"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateAccessToken(user models.User) (string, error)
	GenerateRefreshToken(user models.User) (RefreshToken, error)
	ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error)
	DecryptAccessTokenPayload(tokenString string) (*v1dto.EncryptedPayload, error)
	StoreRefreshToken(ctx context.Context, token RefreshToken) error
	ValidaRefreshToken(ctx context.Context, token string) (RefreshToken, error)
	RevokeRefreshToken(ctx context.Context,token string) error
	ValidTurnstile(token, remoteip string) (*TurnstileResponse, error)
}

type Oauth2Service interface {
	OAuth2Login(provider string) (string, string, error)
	OAuth2Callback(provider, code string) (models.User, error) 
}