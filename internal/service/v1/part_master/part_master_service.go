package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/part_master"
	"github.com/doug-martin/goqu/v9"
)

type partMasterService struct {
	repo repository.PartMasterRepository
}

func NewPartMasterService(repo repository.PartMasterRepository) PartMasterService {
	return &partMasterService{
		repo: repo,
	}
}

func (ss *partMasterService) GetAllPartMasters(params v1dto.GetAllPartMasterParams) ([]models.PartMaster, int64, error) {
	return ss.repo.GetAllPartMasters(params)
}

func (ss *partMasterService) CreatePartMaster(params v1dto.PartMasterParamsInput) error {
	return ss.repo.CreatePartMaster(params)
}

func (ss *partMasterService) EditPartMaster(id int) (models.PartMaster, error) {
	return ss.repo.FindPartMasterById(id)
}

func (ss *partMasterService) UpdatePartMaster(params v1dto.PartMasterParamsUpdate) error {
	updateData := goqu.Record{}
	if params.SkillId != nil {
		updateData["skill_id"] = params.SkillId
	}
	if params.Name != nil {
		updateData["name"] = params.Name
	}
	if params.PartNumber != nil {
		updateData["part_number"] = params.PartNumber
	}
	if params.Status != nil {
		updateData["status"] = params.Status
	}
	
	return ss.repo.UpdatePartMasterById(params.Id, updateData)
}