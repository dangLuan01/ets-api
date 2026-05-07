package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type UserAttemptRepository interface {
	UpsertUserAttempt(attempt models.UserAttempt) (int64, error)
	UpdateUserAttempt(attemptId int64, params models.UserAttempt) error
	UpdateTimeSpentSecAttempt(attemptId int64, TimeSpentSec int) error
	UpsertUserAnswer(answer models.UserAnswer) error
	FindExamResume(userId, examSlug string) (models.ResumeAttempt, bool, error)
	FindAnswerResume(attemptId int) ([]models.ResumeAnswer, error)
	FindAttemptByUserUUID(userUUID string, params v1dto.GetAttemptByUserUuuidParams) ([]v1dto.UserAttemptDTO, int64, error)
}
