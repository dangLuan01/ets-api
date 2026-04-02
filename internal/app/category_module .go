package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/category"
	repository "github.com/dangLuan01/ets-api/internal/repository/category"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/category"
)

type CategoryModule struct {
	routes routes.Route
}

func NewCategoryModule(ctx *ModuleContext) *CategoryModule {

	categoryRepo := repository.NewSqlCategoryRepository(ctx.DB)
	categoryService := v1service.NewCategoryService(categoryRepo)
	categoryHandler := v1handler.NewCategoryHandler(categoryService)
	categoryRoutes := v1routes.NewCategoryRoutes(categoryHandler)

	return &CategoryModule{
		routes: categoryRoutes,
	}
}

func (c *CategoryModule) Routes() routes.Route {
	return c.routes
}