package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/question"
	"github.com/gin-gonic/gin"
)

type QuestionRoutes struct {
	handler *v1handler.QuestionHandler
}

func NewQuestionRoutes(handler *v1handler.QuestionHandler) *QuestionRoutes {
	return &QuestionRoutes {
		handler: handler,
	}
}

func (cr *QuestionRoutes) Register(r *gin.RouterGroup) {
	question := r.Group("/questions")
	{	
		question.POST("/single/create", cr.handler.CreateQuestion)
		question.POST("/group/create", cr.handler.CreateQuestionGroup)
	}
}