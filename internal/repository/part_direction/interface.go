package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type PartDirectionRepository interface {
	FindDirectionByExamId(examId int) ([]models.Direction, error)
	FindDirectionByExamIdAndPartId(examId, partId int) (models.Direction, error)
	CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error
	UpdatePartDirection(params v1dto.UpdatePartDirectionInputParams) error
}