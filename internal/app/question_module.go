package app

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/question"
	repository "github.com/dangLuan01/ets-api/internal/repository/question"
	"github.com/dangLuan01/ets-api/internal/routes"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/question"
)

type QuestionModule struct {
	routes routes.Route
}

func NewQuestionModule(ctx *ModuleContext) *QuestionModule {

	questionRepo := repository.NewSqlQuestionRepository(ctx.DB)
	questionService := v1service.NewQuestionService(questionRepo)
	questionHandler := v1handler.NewQuestionHandler(questionService)
	questionRoutes := v1routes.NewQuestionRoutes(questionHandler)

	return &QuestionModule{
		routes: questionRoutes,
	}
}

func (q *QuestionModule) Routes() routes.Route {
	return q.routes
}