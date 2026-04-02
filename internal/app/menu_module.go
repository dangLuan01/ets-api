package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/menu"
	repository "github.com/dangLuan01/ets-api/internal/repository/menu"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/client"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/menu"
)



type MenuModule struct {
	routes routes.Route
}

func NewMenuModule(ctx *ModuleContext) *MenuModule {

	menuRepo := repository.NewSqlMenuRepository(ctx.DB)
	menuService := v1service.NewMenuService(menuRepo)
	menuHandler := v1handler.NewMenuHandler(menuService)
	menuRoutes := v1routes.NewMenuRoutes(menuHandler)

	return &MenuModule{
		routes: menuRoutes,
	}
}

func (m *MenuModule) Routes() routes.Route {
	return m.routes
}