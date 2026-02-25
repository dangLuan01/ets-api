package repository

import (
	"time"

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
	FindExamById(examId string) (models.Exam, error)
	FindExamQuestionMappingById(examId string) ([]models.ExamQuestionMapping, error)
	FindQuesionByIds(singleIDs []int) ([]models.Question, error)
	FindGroupQuestionByIds(groupIDs []int) ([]models.QuestionGroup, error)
	FindSubQuesionByGroupIds(groupIDs []int) ([]models.Question, error)
}
