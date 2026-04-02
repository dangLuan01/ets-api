package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	repository "github.com/dangLuan01/ets-api/internal/repository/question"
)

type questionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(repo repository.QuestionRepository) QuestionService {
	return &questionService{
		repo: repo,
	}
}

func (qs *questionService) CreateQuestion(params v1dto.QuestionParamsInput) error {
	questionId, err := qs.repo.CreateQuestion(params)
	if err != nil {
		return err
	}

	paramsQuestionMapping := v1dto.ExamQuestionMappingInput{
		ExamId: params.ExamId,
		EntityType: params.EntityType,
		EntityId: questionId,
		OrderIndex: params.SubOrder,
		PartId: params.PartId,
	}

	return qs.repo.CreateQuestionMapping(paramsQuestionMapping)
}

func (qs *questionService) CreateQuestionGroup(params v1dto.QuestionGroupParamsInput) error {
	groupId, err := qs.repo.CreateQuestionGroup(params)
	if err != nil {
		return err
	}

	var paramsQuestion []v1dto.QuestionParamsInput
	for _, subQuestion := range params.SubQuestions {
		paramsQuestion = append(paramsQuestion, v1dto.QuestionParamsInput{
			GroupId: &groupId,
			Part: subQuestion.Part,
			QuestionText: subQuestion.QuestionText,
			ImageUrl: subQuestion.ImageUrl,
			CorrectAnswer: subQuestion.CorrectAnswer,
			OptionA: subQuestion.OptionA,
			OptionB: subQuestion.OptionB,
			OptionC: subQuestion.OptionC,
			OptionD: subQuestion.OptionD,
			AudioStartMs: subQuestion.AudioStartMs,
			AudioEndMs: subQuestion.AudioEndMs,
			SubOrder: subQuestion.SubOrder,
			Explanation: subQuestion.Explanation,
			Transcript: subQuestion.Transcript,
			Tags: subQuestion.Tags,
		})
	}

	err = qs.repo.CreateQuestions(paramsQuestion)
	if err != nil {
		return err
	}

	paramsQuestionGroupMapping := v1dto.ExamQuestionMappingInput{
		ExamId: params.ExamId,
		EntityType: params.EntityType,
		EntityId: groupId,
		OrderIndex: params.SubQuestions[0].SubOrder,
		PartId: params.PartId,
	}

	err = qs.repo.CreateQuestionGroupMapping(paramsQuestionGroupMapping)
	if err != nil {
		return err
	}

	return nil
}