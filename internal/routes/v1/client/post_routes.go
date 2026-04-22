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
	post := r.Group("/client/post")
	{
		post.GET("/get-all", mr.handler.FindAllPosts)
		post.GET("/:slug", mr.handler.FindPostBySlug)
		post.GET("/tag/:slug", mr.handler.FindPostByTagSlug)
	}
}