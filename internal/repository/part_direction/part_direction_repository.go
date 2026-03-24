package repository

import (
	"encoding/json"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_PART_DIRECTION		= "part_directions"
)

type SqlPartDirectionRepository struct {
	db *goqu.Database
}

func NewSqlPartDirectionRepository(DB *goqu.Database) PartDirectionRepository {
	return &SqlPartDirectionRepository{
		db: DB,
	}
}

func (rt *SqlPartDirectionRepository) FindDirectionByExamId(examId int) ([]models.Direction, error) {
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

func (rt *SqlPartDirectionRepository) FindDirectionByExamIdAndPartId(examId, partId int) (models.Direction, error) {
	var partDirection models.Direction

	_, err := rt.db.From(TABLE_PART_DIRECTION).
		Select(
			goqu.C("id"),
			goqu.C("exam_id"),
			goqu.C("direction_text"),
			goqu.C("part_id"),
			goqu.C("audio_start_ms"),
			goqu.C("audio_end_ms"),
			goqu.C("example_data"),
		).
		Where(
			goqu.C("exam_id").Eq(examId),
			goqu.C("part_id").Eq(partId),
		).
		ScanStruct(&partDirection)
	
	if err != nil {
		
		return models.Direction{}, err
	}

	return partDirection, nil
}

func (rt *SqlPartDirectionRepository) CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error {
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

func (rt *SqlPartDirectionRepository) UpdatePartDirection(params v1dto.UpdatePartDirectionInputParams) error {
	jsonBytes, err := json.Marshal(params.ExampleData)
	if err != nil {
		return err
	}
	updateData := goqu.Record{
		"direction_text": params.Direction,
		"audio_start_ms": params.AudioStartMs,
		"audio_end_ms":   params.AudioEndMs,
	}
	if len(params.ExampleData) > 0 {
		updateData["example_data"] = jsonBytes
	} else {
		updateData["example_data"] = nil
	}

	_, err = rt.db.Update(TABLE_PART_DIRECTION).Set(updateData).
		Where(
			goqu.C("exam_id").Eq(params.ExamId),
			goqu.C("part_id").Eq(params.PartId),
		).
		Executor().Exec()

	return err
}