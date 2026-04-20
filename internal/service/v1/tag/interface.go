package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type TagService interface {
	GetAllTags(params v1dto.GetAllTagParams) ([]models.Tag, int64, error)
	CreateTag(params v1dto.TagParamsInput) error
	EditTag(id int) (models.Tag, error)
	UpdateTag(params v1dto.TagParamsUpdate) error
}