package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/user"
	repository "github.com/dangLuan01/ets-api/internal/repository/user"
	repositoryAttempt "github.com/dangLuan01/ets-api/internal/repository/user_attempt"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/client"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/user"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(ctx *ModuleContext) *UserModule {

	userRepo 	:= repository.NewSqlUserRepository(ctx.DB)
	attemptRepo := repositoryAttempt.NewSqlUserAttemptRepository(ctx.DB)
	userService := v1service.NewUserService(userRepo, attemptRepo)
	UserHandler := v1handler.NewUserHandler(userService)
	userRoutes 	:= v1routes.NewUserRoutes(UserHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}