package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type PartMasterRepository interface {
	GetAllPartMasters() ([]models.PartMaster, error)
	CreatePartMaster(params v1dto.PartMasterParamsInput) error
	FindPartMasterById(id int) (models.PartMaster, error)
	UpdatePartMasterById(id int, params goqu.Record) error
}
