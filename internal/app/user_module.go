package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1"
	"github.com/dangLuan01/ets-api/internal/repository"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(ctx *ModuleContext) *UserModule {

	userRepo := repository.NewSqlUserRepository(ctx.DB)
	userService := v1service.NewUserService(userRepo)
	UserHandler := v1handler.NewUserHandler(userService)
	userRoutes := v1routes.NewUserRoutes(UserHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}