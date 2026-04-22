package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/post"
	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	handler *v1handler.PostHandler
}

func NewPostAdminRoutes(handler *v1handler.PostHandler) *PostRoutes {
	return &PostRoutes {
		handler: handler,
	}
}

func (cr *PostRoutes) Register(r *gin.RouterGroup) {
	post := r.Group("/post")
	{
		post.GET("/get-all", cr.handler.GetAllPosts)
		post.POST("/create", cr.handler.CreatePost)
		post.GET("/edit/:id", cr.handler.EditPost)
		post.PUT("/update", cr.handler.UpdatePost)
		post.DELETE("/delete/:id", cr.handler.DeletePost)
	}
}