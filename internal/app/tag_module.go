package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/tag"
	repository "github.com/dangLuan01/ets-api/internal/repository/tag"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/admin"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/tag"
)

type TagModule struct {
	routes routes.Route
}

func NewTagModule(ctx *ModuleContext) *TagModule {

	tagRepo := repository.NewSqlTagRepository(ctx.DB)
	tagService := v1service.NewTagService(tagRepo)
	tagHandler := v1handler.NewTagHandler(tagService)
	tagRoutes := v1routes.NewTagRoutes(tagHandler)

	return &TagModule{
		routes: tagRoutes,
	}
}

func (c *TagModule) Routes() routes.Route {
	return c.routes
}