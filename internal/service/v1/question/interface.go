package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
)

type QuestionService interface {
	CreateQuestion(params v1dto.QuestionParamsInput) error
	CreateQuestionGroup(params v1dto.QuestionGroupParamsInput) error
}