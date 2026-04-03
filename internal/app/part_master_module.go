package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/part_master"
	repository "github.com/dangLuan01/ets-api/internal/repository/part_master"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/admin"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/part_master"
)

type PartMasterModule struct {
	routes routes.Route
}

func NewPartMasterModule(ctx *ModuleContext) *PartMasterModule {

	partMasterRepo := repository.NewSqlPartMasterRepository(ctx.DB)
	partMasterService := v1service.NewPartMasterService(partMasterRepo)
	partMasterHandler := v1handler.NewPartMasterHandler(partMasterService)
	partMasterRoutes := v1routes.NewPartMasterRoutes(partMasterHandler)

	return &PartMasterModule{
		routes: partMasterRoutes,
	}
}

func (c *PartMasterModule) Routes() routes.Route {
	return c.routes
}