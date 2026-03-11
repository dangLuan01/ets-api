package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_SKILLS 			= "skills"
)

type SqlSkillRepository struct {
	db *goqu.Database
}

func NewSqlSkillRepository(DB *goqu.Database) SkillRepository {
	return &SqlSkillRepository{
		db: DB,
	}
}

func (cr *SqlSkillRepository) GetAllSkills() ([]models.SkillMaster, error) {
	var skills []models.SkillMaster
	err := cr.db.From(TABLE_SKILLS).ScanStructs(&skills)
	if err != nil {
		return nil, err
	}
	
	return skills, nil
}

func (cr *SqlSkillRepository) CreateSkill(params v1dto.SkillParamsInput) error {
	_, err := cr.db.From(TABLE_SKILLS).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (cr *SqlSkillRepository) FindSkillById(id int) (skill models.SkillMaster, err error) {

	found, err := cr.db.From(TABLE_SKILLS).Where(goqu.C("id").Eq(id)).ScanStruct(&skill)
	if err != nil {
		return models.SkillMaster{}, err
	}

	if !found {
		return models.SkillMaster{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found skill.")
	}

	return skill, nil
}

func (cr *SqlSkillRepository) UpdateSkillById(id int, params goqu.Record) error {
	
	_, err := cr.db.From(TABLE_SKILLS).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}