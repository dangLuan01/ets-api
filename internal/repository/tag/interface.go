package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type TagRepository interface {
	GetAllTags(params v1dto.GetAllTagParams) ([]models.Tag, int64, error)
	CreateTag(params v1dto.TagParamsInput) error
	FindTagById(id int) (models.Tag, error)
	UpdateTagById(id int, params goqu.Record) error
}