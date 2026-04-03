package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/part_master"
	"github.com/gin-gonic/gin"
)

type PartMasterRoutes struct {
	handler *v1handler.PartMasterHandler
}

func NewPartMasterRoutes(handler *v1handler.PartMasterHandler) *PartMasterRoutes {
	return &PartMasterRoutes {
		handler: handler,
	}
}

func (cr *PartMasterRoutes) Register(r *gin.RouterGroup) {
	partMaster := r.Group("/part-masters")
	{
		partMaster.GET("/get-all", cr.handler.GetAllPartMasters)
		partMaster.POST("/create", cr.handler.CreatePartMaster)
		partMaster.GET("/edit/:id", cr.handler.EditPartMaster)
		partMaster.PUT("/update", cr.handler.UpdatePartMaster)
		partMaster.DELETE("/delete/:id", cr.handler.DeletePartMaster)
	}
}