package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/skill"
	"github.com/dangLuan01/ets-api/internal/repository/skill"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/skill"
)

type SkillModule struct {
	routes routes.Route
}

func NewSkillModule(ctx *ModuleContext) *SkillModule {

	skillRepo := repository.NewSqlSkillRepository(ctx.DB)
	skillService := v1service.NewSkillService(skillRepo)
	skillHandler := v1handler.NewSkillHandler(skillService)
	skillRoutes := v1routes.NewSkillRoutes(skillHandler)

	return &SkillModule{
		routes: skillRoutes,
	}
}

func (c *SkillModule) Routes() routes.Route {
	return c.routes
}