package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/auth"
	repository "github.com/dangLuan01/ets-api/internal/repository/user"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/auth"
	"github.com/dangLuan01/ets-api/pkg/auth"
	"github.com/dangLuan01/ets-api/pkg/cache"
)

type AuthModule struct {
	routes routes.Route
}

func NewAuthModule(ctx *ModuleContext, tokenService auth.TokenService, cacheService cache.RedisCacheService) *AuthModule {

	userRepo := repository.NewSqlUserRepository(ctx.DB)
	authService := v1service.NewAuthService(userRepo, tokenService, cacheService)
	authHandler := v1handler.NewAuthHandler(authService) 
	authRoutes := v1routes.NewAuthRoutes(authHandler)

	return &AuthModule{
		routes: authRoutes,
	}
}

func (m *AuthModule) Routes() routes.Route {
	return m.routes
}