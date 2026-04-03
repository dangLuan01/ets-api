package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/certificate"
	"github.com/gin-gonic/gin"
)

type CertificateRoutes struct {
	handler *v1handler.CertificateHandler
}

func NewCertificateRoutes(handler *v1handler.CertificateHandler) *CertificateRoutes {
	return &CertificateRoutes {
		handler: handler,
	}
}

func (cr *CertificateRoutes) Register(r *gin.RouterGroup) {
	certificate := r.Group("/certificates")
	{
		certificate.GET("/get-all", cr.handler.GetAllCertificates)
		certificate.POST("/create", cr.handler.CreateCertificate)
		certificate.GET("/edit/:id", cr.handler.EditCertificate)
		certificate.PUT("/update", cr.handler.UpdateCertificate)
		certificate.DELETE("/delete/:id", cr.handler.DeleteCertificate)
	}
}