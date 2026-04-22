package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/tag"
	"github.com/gin-gonic/gin"
)

type TagRoutes struct {
	handler *v1handler.TagHandler
}

func NewTagRoutes(handler *v1handler.TagHandler) *TagRoutes {
	return &TagRoutes {
		handler: handler,
	}
}

func (tr *TagRoutes) Register(r *gin.RouterGroup) {
	tag := r.Group("/client/tag")
	{
		tag.GET("/get-all", tr.handler.GetAllTags)
	}
}