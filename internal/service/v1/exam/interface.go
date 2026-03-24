package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type ExamService interface {
	FindExamById(examId int) (models.Exam, error)
	CalculateScoreExam(params v1dto.QuestionAnswerInputParams) (v1dto.DetailExamScore, error)
	// --- SERVICE CHO ADMIN (CRUD EXAM) ---
	GetAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64, error)
	CreateExam(params v1dto.CreateExamInputParams) error
	EditExamById(examId int) (models.ExamModel, error)
	UpdateExam(params v1dto.UpdateExamInputParams) error

	CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error
	UpdatePartDirection(params v1dto.UpdatePartDirectionInputParams) error
	
	GetExamStructure(examId int) (v1dto.ExamStructure, error)
	GetExamPart(examId int, partId int) (v1dto.ExamPart, error)
	UpdateQuestionSingle(params v1dto.UpdateQuestionSingleInputParams) error
	UpdateQuestionGroup(params v1dto.UpdateQuestionGroupInputParams) error
}