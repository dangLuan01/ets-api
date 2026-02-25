package repository

import (
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_EXAM 					= "exams"
	TABLE_EXAM_QUESTION_MAPPING = "exam_question_mappings"
	TABLE_QUESTION_GROUPS 		= "question_groups"
	TABLE_QUESTIONS 			= "questions"
)

type SqlExamRepository struct {
	db *goqu.Database
}

func NewSqlExamRepository(DB *goqu.Database) ExamRepository {
	return &SqlExamRepository{
		db: DB,
	}
}

func (rt *SqlExamRepository) FindExamById(examId string) (models.Exam, error) {
	var exam models.Exam

	found, err := rt.db.From(TABLE_EXAM).
	Select(
		goqu.C("id"),
		goqu.C("title"),
		goqu.C("year"),
		goqu.C("category"),
		goqu.C("total_time"),
		goqu.C("description"),
		goqu.C("thumbnail"),
		goqu.C("audio_full_url"),
		goqu.C("status"),
		goqu.C("created_at"),
	).
	Where(goqu.C("id").Eq(examId)).ScanStruct(&exam)
	
	if !found && err == nil {
		return models.Exam{}, utils.NewError(string(utils.ErrCodeNotFound), "Not found exam.")
	}

	if err != nil {
		return models.Exam{}, err
	}

	return exam, nil
}

func (rt *SqlExamRepository) FindExamQuestionMappingById(examId string) ([]models.ExamQuestionMapping, error) {
	var sections []models.ExamQuestionMapping
	err := rt.db.From(TABLE_EXAM_QUESTION_MAPPING).
		Select(
			goqu.C("id"),
			goqu.C("exam_id"),
			goqu.C("entity_type"),
			goqu.C("entity_id"),
			goqu.C("order_index"),
			goqu.C("part"),
		).
		Where(goqu.C("exam_id").Eq(examId)).
		Order(
			goqu.C("part").Desc(),
			goqu.C("order_index").Asc(),
		).
		ScanStructs(&sections)

	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (rt *SqlExamRepository) FindQuesionByIds(singleIDs []int) ([]models.Question, error) {
	var questions []models.Question

	err := rt.db.From(TABLE_QUESTIONS).
		Select(
			goqu.C("id"),
			goqu.C("group_id"),
			goqu.C("question_text"),
			goqu.C("image_url"),
			goqu.C("audio_start_ms"),
			goqu.C("audio_end_ms"),
			goqu.C("option_a"),
			goqu.C("option_b"),
			goqu.C("option_c"),
			goqu.C("option_d"),
			goqu.C("sub_order"),
			goqu.C("correct_answer"),
			goqu.C("explanation"),
			goqu.C("transcript"),
			goqu.C("difficulty"),
			goqu.C("tags"),
		).
		Where(goqu.C("id").In(singleIDs)).
		ScanStructs(&questions)
	
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (rt *SqlExamRepository) FindGroupQuestionByIds(groupIDs []int) ([]models.QuestionGroup, error) {
	var groupQuestions	[]models.QuestionGroup			

	err := rt.db.From(TABLE_QUESTION_GROUPS).
		Select(
			goqu.C("id"),
			goqu.C("part"),
			goqu.C("passage_text"),
			goqu.C("image_url"),
			goqu.C("audio_start_ms"),
			goqu.C("audio_end_ms"),
			goqu.C("transcript"),
			goqu.C("explanation"),
		).
		Where(goqu.C("id").In(groupIDs)).
		ScanStructs(&groupQuestions)
	
	if err != nil {
		return nil, err
	}

	return groupQuestions, nil
}

func (rt *SqlExamRepository) FindSubQuesionByGroupIds(groupIDs []int) ([]models.Question, error) {
	var subQuestions []models.Question

	err := rt.db.From(TABLE_QUESTIONS).
		Select(
			goqu.C("id"),
			goqu.C("group_id"),
			goqu.C("question_text"),
			goqu.C("image_url"),
			goqu.C("audio_start_ms"),
			goqu.C("audio_end_ms"),
			goqu.C("option_a"),
			goqu.C("option_b"),
			goqu.C("option_c"),
			goqu.C("option_d"),
			goqu.C("sub_order"),
			goqu.C("correct_answer"),
			goqu.C("explanation"),
			goqu.C("transcript"),
			goqu.C("difficulty"),
			goqu.C("tags"),
		).
		Where(goqu.C("group_id").In(groupIDs)).
		Order(goqu.C("sub_order").Asc()).
		ScanStructs(&subQuestions)
	
	if err != nil {
		return nil, err
	}

	return subQuestions, nil
}