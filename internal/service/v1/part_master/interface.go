package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type PartMasterService interface {
	GetAllPartMasters() ([]models.PartMaster, error)
	CreatePartMaster(params v1dto.PartMasterParamsInput) error
	EditPartMaster(id int) (models.PartMaster, error)
	UpdatePartMaster(params v1dto.PartMasterParamsUpdate) error
}