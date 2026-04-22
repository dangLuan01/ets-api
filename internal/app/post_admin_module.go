package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/post"
	repository "github.com/dangLuan01/ets-api/internal/repository/post"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/admin"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/post"
)

type PostAdminModule struct {
	routes routes.Route
}

func NewPostAdminModule(ctx *ModuleContext) *PostAdminModule {

	postRepo := repository.NewSqlPostRepository(ctx.DB)
	postService := v1service.NewPostService(postRepo, ctx.DB)
	postHandler := v1handler.NewPostHandler(postService)
	postRoutes := v1routes.NewPostAdminRoutes(postHandler)

	return &PostAdminModule{
		routes: postRoutes,
	}
}

func (c *PostAdminModule) Routes() routes.Route {
	return c.routes
}