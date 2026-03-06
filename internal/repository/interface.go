package repository

import (
	"time"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindBYUUID(uuid string) (models.User, error)
	Create(user models.User) error
	Update(uuid uuid.UUID, user models.User) error
	Delete(uuid uuid.UUID) error
	FindByEmail(email string) (models.User, error)
	UpdateMember(uuid string, is_member int, expriedDate time.Time) error
	UpdatePassword(uuid string, password string) error
	UpdateCountUpload(uuid string) error
}

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
}
