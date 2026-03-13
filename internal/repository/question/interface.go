package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
)

type QuestionRepository interface {
	CreateQuestion(paramsQuestion v1dto.QuestionParamsInput) (int64, error)
	CreateQuestionMapping(params v1dto.ExamQuestionMappingInput) error
	CreateQuestionGroupMapping(params v1dto.ExamQuestionMappingInput) error
	CreateQuestions(params []v1dto.QuestionParamsInput) (error)
	CreateQuestionGroup(params v1dto.QuestionGroupParamsInput) (int64, error)
}
