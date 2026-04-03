package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/certificate"
	repository "github.com/dangLuan01/ets-api/internal/repository/certificate"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/admin"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/certificate"
)

type CertificateModule struct {
	routes routes.Route
}

func NewCertificateModule(ctx *ModuleContext) *CertificateModule {

	certificateRepo := repository.NewSqlCertificateRepository(ctx.DB)
	certificateService := v1service.NewCertificateService(certificateRepo)
	certificateHandler := v1handler.NewCertificateHandler(certificateService)
	certificateRoutes := v1routes.NewCertificateRoutes(certificateHandler)

	return &CertificateModule{
		routes: certificateRoutes,
	}
}

func (c *CertificateModule) Routes() routes.Route {
	return c.routes
}