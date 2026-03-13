package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type SkillService interface {
	GetAllSkills(params v1dto.GetAllSkillParams) ([]models.SkillMaster, int64, error)
	CreateSkill(params v1dto.SkillParamsInput) error
	EditSkill(id int) (models.SkillMaster, error)
	UpdateSkill(params v1dto.SkillParamsUpdate) error
}