package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/exam"
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
		exam.POST("/calculate/score", tr.handler.CalculateScoreExam)
		//ROUTE FOR CRUD EXAM (ADMIN)
		exam.GET("/get-all", tr.handler.GetAllExams)
		exam.POST("/create", tr.handler.CreateExam)
		exam.GET("/edit/:id", tr.handler.EditExam)
		exam.PUT("/update", tr.handler.UpdateExam)
		//Route for CRUD Advance (ADMIN)
		//exam.GET("/:id/structure", tr.handler.GetExamStructure)
		exam.POST("/part-direction/create", tr.handler.CreatePartDirection)
	}
}