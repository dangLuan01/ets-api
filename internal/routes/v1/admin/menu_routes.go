package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/menu"
	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	handler *v1handler.MenuHandler
}

func NewMenuAdminRoutes(handler *v1handler.MenuHandler) *MenuRoutes {
	return &MenuRoutes {
		handler: handler,
	}
}

func (cr *MenuRoutes) Register(r *gin.RouterGroup) {
	menu := r.Group("/menu")
	{
		menu.GET("/get-all", cr.handler.GetAllMenu)
		menu.GET("/structure", cr.handler.GetMenuStructure)
		menu.POST("/create", cr.handler.CreateMenu)
		menu.GET("/edit/:id", cr.handler.EditMenu)
		menu.PUT("/update", cr.handler.UpdateMenu)
	}
}