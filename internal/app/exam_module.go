package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/exam"
	"github.com/dangLuan01/ets-api/internal/repository/exam"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/exam"
)

type ExamModule struct {
	routes routes.Route
}

func NewExamModule(ctx *ModuleContext) *ExamModule {

	examRepo := repository.NewSqlExamRepository(ctx.DB)
	examService := v1service.NewExamService(examRepo)
	examHandler := v1handler.NewExamHandler(examService)
	examRoutes := v1routes.NewExamRoutes(examHandler)

	return &ExamModule{
		routes: examRoutes,
	}
}

func (e *ExamModule) Routes() routes.Route {
	return e.routes
}