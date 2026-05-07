package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_USER_ATTEMPT	= "user_attempts"
	TABLE_USER_ANSWERS	= "user_answers"
	TABLE_EXAM 			= "exams"
)

type UpdateUserAttempt struct {
	EndTime 		*string `db:"end_time"`
	ListeningScore 	int 	`db:"listening_score"`
	ReadingScore 	int 	`db:"reading_score"`
	TotalScore		int		`db:"total_score"`
	Status			int8	`db:"status"`
}

type SqlUserAttemptRepository struct {
	db *goqu.Database
}

func NewSqlUserAttemptRepository(DB *goqu.Database) UserAttemptRepository {
	return &SqlUserAttemptRepository{
		db: DB,
	}
}

func (rt *SqlUserAttemptRepository) UpsertUserAttempt(attempt models.UserAttempt) (int64, error) {
	resp, err := rt.db.Insert(TABLE_USER_ATTEMPT).Rows(attempt).
		OnConflict(
			goqu.DoUpdate("user_id,exam_slug", goqu.Record{
				"start_time": attempt.StartTime,
				"end_time": nil,
				"listening_score": 0,
				"reading_score": 0,
				"total_score": 0,
				"status": 1,
			}),
		).Executor().Exec()
	if err != nil{
    	return 0, err
	}

	attemptId, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	if _, err := rt.db.Delete(TABLE_USER_ANSWERS).
		Where(goqu.C("attempt_id").Eq(attemptId)).
		Executor().Exec(); err != nil {
			return 0, err
		}
	
	return attemptId, nil
}

func (rt *SqlUserAttemptRepository) UpdateUserAttempt(attemptId int64, params models.UserAttempt) error {
	_, err := rt.db.Update(TABLE_USER_ATTEMPT).Where(
		goqu.C("id").Eq(attemptId),
	).Set(
		UpdateUserAttempt{
			EndTime: params.EndTime,
			ListeningScore: params.ListeningScore,
			ReadingScore: params.ReadingScore,
			TotalScore: params.TotalScore,
			Status: params.Status,
		},
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}

func (rt *SqlUserAttemptRepository) UpdateTimeSpentSecAttempt(attemptId int64, TimeSpentSec int) error {
	_, err := rt.db.Update(TABLE_USER_ATTEMPT).Where(
		goqu.C("id").Eq(attemptId),
	).Set(
		goqu.Record{
			"time_spent_sec": TimeSpentSec,
		},
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}

func (rt *SqlUserAttemptRepository) UpsertUserAnswer(answer models.UserAnswer) error {
	_, err := rt.db.Insert(TABLE_USER_ANSWERS).Rows(answer).
		OnConflict(
			goqu.DoUpdate("attempt_id,question_id", goqu.Record{
				"selected_answer": answer.SelectedAnswer,
			}),
		).
		Executor().Exec()
	if err != nil{
    	return err
	}

	return nil
}

func (rt *SqlUserAttemptRepository) FindExamResume(userId, examSlug string) (models.ResumeAttempt, bool, error) {
	var attempt models.ResumeAttempt
	found, err := rt.db.From(TABLE_USER_ATTEMPT).
		Select(
			goqu.C("id"),
			goqu.C("user_id"),
			goqu.C("exam_slug"),
			goqu.C("status"),
			goqu.C("time_spent_sec"),
		).Where(
			goqu.C("user_id").Eq(userId),
			goqu.C("exam_slug").Eq(examSlug),
		).ScanStruct(&attempt)

	if !found {
		return models.ResumeAttempt{}, false, nil
	}

	if err != nil {
		return models.ResumeAttempt{}, false, err
	}

	return attempt, true, nil
}

func (rt *SqlUserAttemptRepository) FindAnswerResume(attemptId int) ([]models.ResumeAnswer, error) {
	var answers []models.ResumeAnswer
	err := rt.db.From(TABLE_USER_ANSWERS).
		Select(
			goqu.C("question_id"),
			goqu.C("selected_answer"),
		).Where(
			goqu.C("attempt_id").Eq(attemptId),
		).ScanStructs(&answers)

	if err != nil {
		return nil, err
	}

	return answers, nil
}

func (rt *SqlUserAttemptRepository) FindAttemptByUserUUID(userUUID string, params v1dto.GetAttemptByUserUuuidParams) ([]v1dto.UserAttemptDTO, int64, error) {
	var attempts []v1dto.UserAttemptDTO
	answerCount := rt.db.From(goqu.T(TABLE_USER_ANSWERS).As("ans")).
		Select(goqu.COUNT("*")).
		Where(
			goqu.I("ans.attempt_id").Eq(goqu.I("ua.id")),
		)

	ds := rt.db.From(goqu.T(TABLE_USER_ATTEMPT).As("ua")).
		InnerJoin(
			goqu.T(TABLE_EXAM).As("e"), goqu.On(
				goqu.I("e.slug").Eq(goqu.I("ua.exam_slug")),
			),
		).
		Select(
			goqu.I("ua.exam_slug"),
			goqu.I("e.total_question"),
			goqu.I("ua.start_time"),
			goqu.I("ua.end_time"),
			goqu.I("ua.listening_score"),
			goqu.I("ua.reading_score"),
			goqu.I("ua.total_score"),
			goqu.I("ua.time_spent_sec"),
			goqu.I("ua.status"),
			goqu.Case().When(
				goqu.I("ua.status").Eq(1),
				answerCount,
			).Else(0).As("total_answer"),
		).Where(
			goqu.I("ua.user_id").Eq(userUUID),
		)

	if params.Status != 0 {
		ds = ds.Where(goqu.I("ua.status").Eq(params.Status))
	}

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}

	if err := ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).ScanStructs(&attempts); err != nil {
		return nil, 0, err
	}

	return attempts, totalRecords, nil
}