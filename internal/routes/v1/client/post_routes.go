package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/post"
	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	handler *v1handler.PostHandler
}

func NewPostRoutes(handler *v1handler.PostHandler) *PostRoutes {
	return &PostRoutes {
		handler: handler,
	}
}

func (mr *PostRoutes) Register(r *gin.RouterGroup) {
	menus := r.Group("/client/post")
	{
		menus.GET("/get-all", mr.handler.FindAllPosts)
		menus.GET("/:slug", mr.handler.FindPostBySlug)
	}
}