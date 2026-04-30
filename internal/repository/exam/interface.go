package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type ExamRepository interface {
	// --- REPO CLIENT ---
	FindExamBySlug(examSlug string) (models.Exam, error)
	FindExamQuestionMappingById(examId int) ([]models.ExamQuestionMapping, error)
	FindQuesionByIds(singleIDs []int) ([]models.Question, error)
	FindGroupQuestionByIds(groupIDs []int) ([]models.QuestionGroup, error)
	FindSubQuesionByGroupIds(groupIDs []int) ([]models.Question, error)
	FindSkillsByCertId(certId int) ([]models.SkillMaster, error)
	FindPartsByCertId(certId int) ([]models.PartMaster, error)
	GetCorrectAnswersWithSkillContext(examId int, ids []int) ([]v1dto.QuestionWithSkill, error)
	GetScoreConversionTable(certId int) ([]models.ScoreConversion, error)
	SaveAttemptWithAnswers(attempt models.UserAttempt, answers []models.UserAnswer) error
	FindFilterStructure()([]*v1dto.FilterStructure, error)
	FindExamsByFilter(params v1dto.FilterExamParams) ([]v1dto.ExamFilterDTO, int64, error)
	FindFeaturedExams(params v1dto.ExamFeaturedParams) ([]v1dto.ExamFeaturedRaw, int64, error)
	// --- REPO ADMIN (CRUD EXAM) ---
	FindExamById(examId int) (models.Exam, error)
	FindAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64, error)
	CreateExam(exam v1dto.CreateExamInputParams) error
	GetExamById(examId int) (models.ExamModel, error)
	//UpdateExam(examId int, params v1dto.UpdateExamInputParams) error
	UpdateExam(tx *goqu.TxDatabase, examId int, data goqu.Record) error
	DeleteExamCategories(tx *goqu.TxDatabase, examId int) error
	InsertExamCategories(tx *goqu.TxDatabase, rows []goqu.Record) error

	DeleteExamQuestions(tx *goqu.TxDatabase, examId int) error
	GetTargetExamQuestions(targetExamId int, partId []int) ([]models.TargetExamMapping, error)
	InsertExamQuestions(tx *goqu.TxDatabase, rows []models.TargetExamMapping) error

	FindExamQuestionMappingByPartId(examId int, partId int) ([]v1dto.ExamQuestionMappingDTO, error)
	UpdateQuestionSingle(params v1dto.UpdateQuestionSingleInputParams) error
	UpdateQuestionGroup(params v1dto.UpdateQuestionGroupInputParams) error
}
