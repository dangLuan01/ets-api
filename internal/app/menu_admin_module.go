package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/menu"
	repository "github.com/dangLuan01/ets-api/internal/repository/menu"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/admin"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/menu"
)

type MenuAdminModule struct {
	routes routes.Route
}

func NewMenuAdminModule(ctx *ModuleContext) *MenuAdminModule {

	menuRepo := repository.NewSqlMenuRepository(ctx.DB)
	menuService := v1service.NewMenuService(menuRepo)
	menuHandler := v1handler.NewMenuHandler(menuService)
	menuRoutes := v1routes.NewMenuAdminRoutes(menuHandler)

	return &MenuAdminModule{
		routes: menuRoutes,
	}
}

func (m *MenuAdminModule) Routes() routes.Route {
	return m.routes
}