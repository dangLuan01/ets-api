package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/doug-martin/goqu/v9"
)
const (
	TABLE_QUESTION				= "questions"
	TABLE_QUESTION_GROUP		= "question_groups"
	TABLE_EXAM_QUESTION_MAPPING	= "exam_question_mappings"
)

type SqlQuestionRepository struct {
	db *goqu.Database
}

func NewSqlQuestionRepository(DB *goqu.Database) QuestionRepository {
	return &SqlQuestionRepository{
		db: DB,
	}
}

func (qr *SqlQuestionRepository) CreateQuestion(params v1dto.QuestionParamsInput) (int64, error) {
	resp, err := qr.db.From(TABLE_QUESTION).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return 0, err
	}

	questionId, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return questionId, nil
}

func (qr *SqlQuestionRepository) CreateQuestionMapping(params v1dto.ExamQuestionMappingInput) error {
	_, err := qr.db.From(TABLE_EXAM_QUESTION_MAPPING).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (qr *SqlQuestionRepository) CreateQuestions(paramsQuestion []v1dto.QuestionParamsInput) (error) {
	_, err := qr.db.From(TABLE_QUESTION).Insert().Rows(paramsQuestion).Executor().Exec()
	if err != nil {
		return err
	}

	return err
}

func (qr *SqlQuestionRepository) CreateQuestionGroup(params v1dto.QuestionGroupParamsInput) (int64, error) {
	resp, err := qr.db.From(TABLE_QUESTION_GROUP).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return 0, err
	}

	groupId, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return groupId, nil
}

func (qr *SqlQuestionRepository) CreateQuestionGroupMapping(params v1dto.ExamQuestionMappingInput) error {
	_, err := qr.db.From(TABLE_EXAM_QUESTION_MAPPING).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}