package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/tag"
	"github.com/gin-gonic/gin"
)

type TagRoutes struct {
	handler *v1handler.TagHandler
}

func NewTagAdminRoutes(handler *v1handler.TagHandler) *TagRoutes {
	return &TagRoutes {
		handler: handler,
	}
}

func (cr *TagRoutes) Register(r *gin.RouterGroup) {
	tag := r.Group("/tag")
	{
		tag.GET("/get-all", cr.handler.GetAllTags)
		tag.POST("/create", cr.handler.CreateTag)
		tag.GET("/edit/:id", cr.handler.EditTag)
		tag.PUT("/update", cr.handler.UpdateTag)
		tag.DELETE("/delete/:id", cr.handler.DeleteTag)
	}
}