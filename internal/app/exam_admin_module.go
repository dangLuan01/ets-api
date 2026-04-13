package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/exam"
	repositoryExam "github.com/dangLuan01/ets-api/internal/repository/exam"
	repositoryPartDirection "github.com/dangLuan01/ets-api/internal/repository/part_direction"
	repositoryQuestion "github.com/dangLuan01/ets-api/internal/repository/question"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1/admin"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/exam"
)

type ExamAdminModule struct {
	routes routes.Route
}

func NewExamAdminModule(ctx *ModuleContext) *ExamAdminModule {
	partDirectionRepo := repositoryPartDirection.NewSqlPartDirectionRepository(ctx.DB)
	questionRepository := repositoryQuestion.NewSqlQuestionRepository(ctx.DB)
	examRepo := repositoryExam.NewSqlExamRepository(ctx.DB)
	examService := v1service.NewExamService(examRepo, ctx.DB, partDirectionRepo, questionRepository)
	examHandler := v1handler.NewExamHandler(examService)
	examRoutes := v1routes.NewExamRoutes(examHandler)

	return &ExamAdminModule{
		routes: examRoutes,
	}
}

func (e *ExamAdminModule) Routes() routes.Route {
	return e.routes
}