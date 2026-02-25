package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type ExamRoutes struct {
	handler *v1handler.ExamHandler
}

func NewExamRoutes(handler *v1handler.ExamHandler) *ExamRoutes {
	return &ExamRoutes {
		handler: handler,
	}
}

func (tr *ExamRoutes) Register(r *gin.RouterGroup) {
	exam := r.Group("/exams")
	{
		exam.POST("/:id/full-test", tr.handler.FindExamById)
	}
}