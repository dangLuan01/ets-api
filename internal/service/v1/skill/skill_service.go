package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/skill"
	"github.com/doug-martin/goqu/v9"
)

type skillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) SkillService {
	return &skillService{
		repo: repo,
	}
}

func (ss *skillService) GetAllSkills() ([]models.SkillMaster, error) {
	return ss.repo.GetAllSkills()
}

func (ss *skillService) CreateSkill(params v1dto.SkillParamsInput) error {
	return ss.repo.CreateSkill(params)
}

func (ss *skillService) EditSkill(id int) (models.SkillMaster, error) {
	return ss.repo.FindSkillById(id)
}

func (ss *skillService) UpdateSkill(params v1dto.SkillParamsUpdate) error {
	updateData := goqu.Record{}
	if params.CertId != nil {
		updateData["cert_id"] = params.CertId
	}
	if params.Code != nil {
		updateData["code"] = params.Code
	}
	if params.Name != nil {
		updateData["name"] = params.Name
	}
	if params.OrderIndex != nil {
		updateData["order_index"] = params.OrderIndex
	}
	if params.Status != nil {
		updateData["status"] = params.Status
	}
	
	return ss.repo.UpdateSkillById(params.Id, updateData)
}