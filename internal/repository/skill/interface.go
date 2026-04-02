package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type SkillRepository interface {
	GetAllSkills(params v1dto.GetAllSkillParams) ([]models.SkillMaster, int64, error)
	CreateSkill(params v1dto.SkillParamsInput) error
	FindSkillById(id int) (models.SkillMaster, error)
	UpdateSkillById(id int, params goqu.Record) error
}