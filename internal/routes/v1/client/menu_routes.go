package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/menu"
	"github.com/gin-gonic/gin"
)


type MenuRoutes struct {
	handler *v1handler.MenuHandler
}

func NewMenuRoutes(handler *v1handler.MenuHandler) *MenuRoutes {
	return &MenuRoutes {
		handler: handler,
	}
}

func (mr *MenuRoutes) Register(r *gin.RouterGroup) {
	menus := r.Group("/menus")
	{
		menus.GET("/", mr.handler.GetMenuHeader)
	}
}