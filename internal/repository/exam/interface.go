package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type ExamRepository interface {
	FindExamById(examId int) (models.Exam, error)
	FindExamQuestionMappingById(examId int) ([]models.ExamQuestionMapping, error)
	FindQuesionByIds(singleIDs []int) ([]models.Question, error)
	FindGroupQuestionByIds(groupIDs []int) ([]models.QuestionGroup, error)
	FindSubQuesionByGroupIds(groupIDs []int) ([]models.Question, error)
	FindSkillsByCertId(certId int) ([]models.SkillMaster, error)
	FindPartsByCertId(certId int) ([]models.PartMaster, error)
	GetCorrectAnswersWithSkillContext(examId int, ids []int) ([]v1dto.QuestionWithSkill, error)
	GetScoreConversionTable(certId int) ([]models.ScoreConversion, error)
	SaveAttemptWithAnswers(attempt models.UserAttempt, answers []models.UserAnswer) error

	// --- REPO CHO ADMIN (CRUD EXAM) ---
	FindAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64, error)
	CreateExam(exam v1dto.CreateExamInputParams) error
	GetExamById(examId int) (models.ExamModel, error)
	UpdateExam(examId int, params goqu.Record) error
	FindExamQuestionMappingByPartId(examId int, partId int) ([]v1dto.ExamQuestionMappingDTO, error)
	UpdateQuestionSingle(params v1dto.UpdateQuestionSingleInputParams) error
	UpdateQuestionGroup(params v1dto.UpdateQuestionGroupInputParams) error
}
