package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type ExamService interface {
	FindExamById(examId int) (models.Exam, error)
	CalculateScoreExam(params v1dto.QuestionAnswerInputParams) (v1dto.DetailExamScore, error)
	// --- SERVICE CHO ADMIN (CRUD EXAM) ---
	GetAllExams() ([]models.ExamModel, error)
	CreateExam(params v1dto.CreateExamInputParams) error
	EditExamById(examId int) (models.ExamModel, error)
	UpdateExam(params v1dto.UpdateExamInputParams) error
}