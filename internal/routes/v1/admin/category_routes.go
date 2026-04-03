package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/category"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	handler *v1handler.CategoryHandler
}

func NewCategoryRoutes(handler *v1handler.CategoryHandler) *CategoryRoutes {
	return &CategoryRoutes {
		handler: handler,
	}
}

func (cr *CategoryRoutes) Register(r *gin.RouterGroup) {
	category := r.Group("/category")
	{
		category.GET("/get-all", cr.handler.GetAllCategory)
		category.GET("/structure", cr.handler.GetCategoryStructure)
		category.POST("/create", cr.handler.CreateCategory)
		category.GET("/edit/:id", cr.handler.EditCategory)
		category.PUT("/update", cr.handler.UpdateCategory)
	}
}