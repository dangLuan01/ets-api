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
	FindDirectionByExamId(examId int) ([]models.Direction, error)
	FindSkillsByCertId(certId int) ([]models.SkillMaster, error)
	FindPartsByCertId(certId int) ([]models.PartMaster, error)
	GetCorrectAnswersWithSkillContext(examId int, ids []int) ([]v1dto.QuestionWithSkill, error)
	GetScoreConversionTable(certId int) ([]models.ScoreConversion, error)
	SaveAttemptWithAnswers(attempt models.UserAttempt, answers []models.UserAnswer) error

	// --- REPO CHO ADMIN (CRUD EXAM) ---
	FindAllExams() ([]models.ExamModel, error)
	CreateExam(exam v1dto.CreateExamInputParams) error
	GetExamById(examId int) (models.ExamModel, error)
	UpdateExam(examId int, params goqu.Record) error
}
