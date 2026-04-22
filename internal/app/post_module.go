package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/post"
	repository "github.com/dangLuan01/ets-api/internal/repository/post"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/client"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/post"
)

type PostModule struct {
	routes routes.Route
}

func NewPostModule(ctx *ModuleContext) *PostModule {

	postRepo := repository.NewSqlPostRepository(ctx.DB)
	postService := v1service.NewPostService(postRepo, ctx.DB)
	postHandler := v1handler.NewPostHandler(postService)
	postRoutes := v1routes.NewPostRoutes(postHandler)

	return &PostModule{
		routes: postRoutes,
	}
}

func (c *PostModule) Routes() routes.Route {
	return c.routes
}