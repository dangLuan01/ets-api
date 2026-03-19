package repository

import (
	"encoding/json"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_EXAM 					= "exams"
	TABLE_CERTIFICATE 			= "certificates"
	TABLE_EXAM_QUESTION_MAPPING = "exam_question_mappings"
	TABLE_QUESTION_GROUP		= "question_groups"
	TABLE_QUESTIONS 			= "questions"
	TABLE_PART_DIRECTION		= "part_directions"
	TABLE_SKILLS				= "skills"
	TABLE_PART_MASTER			= "part_masters"
	TABLE_SCORE_CONVERSION		= "score_conversion_tables"
	TABLE_USER_ATTEMPT			= "user_attempts"
	TABLE_USER_ANSWERS			= "user_answers"
)


type SqlExamRepository struct {
	db *goqu.Database
}

func NewSqlExamRepository(DB *goqu.Database) ExamRepository {
	return &SqlExamRepository{
		db: DB,
	}
}

func (rt *SqlExamRepository) FindExamById(examId int) (models.Exam, error) {
	var exam models.Exam

	found, err := rt.db.From(goqu.T(TABLE_EXAM).As("e")).
	Select(
		goqu.I("e.id"),
		goqu.I("e.cert_id"),
		goqu.I("e.title"),
		goqu.I("e.year"),
		goqu.I("e.category"),
		goqu.I("e.total_time"),
		goqu.I("e.total_question"),
		goqu.I("e.description"),
		goqu.I("e.thumbnail"),
		goqu.I("e.audio_full_url"),
		goqu.I("e.status"),
		goqu.I("e.created_at"),
		goqu.I("s.code").As("cert_code"),
	).
	Join(goqu.T(TABLE_CERTIFICATE).As("s"), goqu.On(goqu.I("s.id").Eq(goqu.I("e.cert_id")))).
	Where(goqu.I("e.id").Eq(examId)).ScanStruct(&exam)
	
	if !found && err == nil {
		return models.Exam{}, utils.NewError(string(utils.ErrCodeNotFound), "Not found exam.")
	}

	if err != nil {
		
		return models.Exam{}, err
	}

	return exam, nil
}

func (rt *SqlExamRepository) FindExamQuestionMappingById(examId int) ([]models.ExamQuestionMapping, error) {
	var sections []models.ExamQuestionMapping
	err := rt.db.From(TABLE_EXAM_QUESTION_MAPPING).
		Select(
			goqu.C("id"),
			goqu.C("exam_id"),
			goqu.C("entity_type"),
			goqu.C("entity_id"),
			goqu.C("order_index"),
			goqu.C("part_id"),
		).
		Where(goqu.C("exam_id").Eq(examId)).
		Order(
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

	err := rt.db.From(TABLE_QUESTION_GROUP).
		Select(
			goqu.C("id"),
			goqu.C("part_id"),
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

func (rt *SqlExamRepository) FindDirectionByExamId(examId int) ([]models.Direction, error) {
	var partDirections []models.Direction

	err := rt.db.From(TABLE_PART_DIRECTION).
		Select(
			goqu.C("id"),
			goqu.C("exam_id"),
			goqu.C("direction_text"),
			goqu.C("part_id"),
			goqu.C("audio_start_ms"),
			goqu.C("audio_end_ms"),
			goqu.C("example_data"),
		).
		Where(goqu.C("exam_id").Eq(examId)).
		Order(goqu.C("part_id").Asc()).
		ScanStructs(&partDirections)
	
	if err != nil {
		
		return nil, err
	}

	return partDirections, nil
}

func (rt *SqlExamRepository) FindSkillsByCertId(certId int) ([]models.SkillMaster, error) {
	var skillsMaster []models.SkillMaster
	err := rt.db.From(TABLE_SKILLS).
		Select(
			goqu.C("id"),
			goqu.C("cert_id"),
			goqu.C("code"),
			goqu.C("name"),
			goqu.C("order_index"),
		).
		Where(goqu.C("cert_id").Eq(certId)).
		Order(goqu.C("order_index").Asc()).
		ScanStructs(&skillsMaster)
	
	if err != nil {

		return nil, err
	}

	return skillsMaster, nil
}

func (rt *SqlExamRepository) FindPartsByCertId(certId int) ([]models.PartMaster, error) {
    var partsMaster []models.PartMaster
    
    err := rt.db.From(goqu.T(TABLE_PART_MASTER).As("pm")).
        Join(
            goqu.T("skills").As("s"),
            goqu.On(goqu.I("pm.skill_id").Eq(goqu.I("s.id"))), // Điều kiện JOIN
        ).
        Select(
            goqu.I("pm.id"),
            goqu.I("pm.skill_id"),
            goqu.I("pm.part_number"),
            goqu.I("pm.name"),
        ).
        Where(goqu.I("s.cert_id").Eq(certId)).
        Order(goqu.I("pm.part_number").Asc()).
        ScanStructs(&partsMaster)
    
    if err != nil {

        return nil, err
    }

    return partsMaster, nil
}

func (rt *SqlExamRepository) GetCorrectAnswersWithSkillContext(examId int, questionIds []int) ([]v1dto.QuestionWithSkill, error) {
	var results []v1dto.QuestionWithSkill

	err := rt.db.From(goqu.T(TABLE_QUESTIONS).As("q")).
		Select(
			goqu.I("q.id").As("question_id"),
			goqu.I("q.correct_answer"),
			goqu.I("pm.skill_id"),
		).
		Join(goqu.T(TABLE_EXAM_QUESTION_MAPPING).As("eqm"), goqu.On(
			goqu.Or(
				goqu.And(
					goqu.I("eqm.entity_type").Eq("SINGLE"),
					goqu.I("eqm.entity_id").Eq(goqu.I("q.id")),
				),
				goqu.And(
					goqu.I("eqm.entity_type").Eq("GROUP"),
					goqu.I("eqm.entity_id").Eq(goqu.I("q.group_id")),
				),
			),
		)).
		Join(goqu.T(TABLE_PART_MASTER).As("pm"), goqu.On(
			goqu.I("eqm.part_id").Eq(goqu.I("pm.id")),
		)).
		Where(
			goqu.I("eqm.exam_id").Eq(examId),
			goqu.I("q.id").In(questionIds),
		).
		ScanStructs(&results)

	return results, err
}

func (rt *SqlExamRepository) GetScoreConversionTable(certId int) ([]models.ScoreConversion, error) {
	var results []models.ScoreConversion

	err := rt.db.From(TABLE_SCORE_CONVERSION).
		Select(
			goqu.C("skill_id"),
			goqu.C("raw_score"),
			goqu.C("scaled_score"),
		).
		Where(goqu.C("cert_id").Eq(certId)).
		ScanStructs(&results)

	return results, err
}

func (rt *SqlExamRepository) SaveAttemptWithAnswers(attempt models.UserAttempt, answers []models.UserAnswer) error {
	tx, err := rt.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	resp, err := rt.db.Insert(TABLE_USER_ATTEMPT).Rows(attempt).
		Executor().Exec()
	if err != nil{
    	return err
	}

	attemptId, err := resp.LastInsertId()
	if err != nil {
		return err
	}

	if len(answers) == 0 {
		return nil
	}

	rows := make([]map[string]interface{}, len(answers))
	for i, ans := range answers {
		rows[i] = map[string]interface{}{
			"attempt_id": attemptId,
			"question_id": ans.QuestionId,
			"selected_answer": ans.SelectedAnswer,
			"is_correct": ans.IsCorrect,
		}
	}

	_, err = tx.Insert(TABLE_USER_ANSWERS).Rows(rows).Executor().Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rt *SqlExamRepository) FindAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64, error) {
	var exams []models.ExamModel
	ds := rt.db.From(TABLE_EXAM)

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}

	if err := ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).ScanStructs(&exams); err != nil {
		return nil, 0, err
	}

	return exams, totalRecords, nil
}

func (rt *SqlExamRepository) CreateExam(exam v1dto.CreateExamInputParams) error {
	_, err := rt.db.Insert(TABLE_EXAM).Rows(exam).Executor().Exec()
	return err
}

func (rt *SqlExamRepository) GetExamById(examId int) (models.ExamModel, error) {
	var exam models.ExamModel
	found, err := rt.db.From(TABLE_EXAM).
		Where(goqu.C("id").Eq(examId)).
		ScanStruct(&exam)
	if err != nil {
		return models.ExamModel{}, err
	}
	if !found {
		return models.ExamModel{}, utils.NewError(string(utils.ErrCodeNotFound), "Not found exam.")
	}

	return exam, nil
}

func (rt *SqlExamRepository) UpdateExam(examId int, params goqu.Record) error {
	_, err := rt.db.From(TABLE_EXAM).
		Update().Set(params).
		Where(goqu.C("id").Eq(examId)).
		Executor().Exec()
		
	return err
}

func (rt *SqlExamRepository) CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error {
	jsonBytes, err := json.Marshal(params.ExampleData)
    if err != nil {
        return err
    }

    insertData := goqu.Record{
        "exam_id":        params.ExamId,
        "part_id":        params.PartId,
        "direction_text": params.Direction,
        "audio_start_ms": params.AudioStartMs,
        "audio_end_ms":   params.AudioEndMs,
    }
	if len(params.ExampleData) > 0 {
        insertData["example_data"] = jsonBytes
    }
	
	_, err = rt.db.Insert(TABLE_PART_DIRECTION).Rows(insertData).Executor().Exec()

	return err
}

func (rt *SqlExamRepository) FindExamQuestionMappingByPartId(examId int, partId int) ([]v1dto.ExamQuestionMappingDTO, error) {
	var mappings []v1dto.ExamQuestionMappingDTO
	err := rt.db.From(TABLE_EXAM_QUESTION_MAPPING).
		Select(
			goqu.C("entity_type"),
			goqu.C("entity_id"),
			goqu.C("order_index"),
		).
		Where(
			goqu.C("exam_id").Eq(examId),
			goqu.C("part_id").Eq(partId),
		).
		Order(goqu.C("order_index").Asc()).
		ScanStructs(&mappings)

	if err != nil {
		return nil, err
	}

	return mappings, nil
}