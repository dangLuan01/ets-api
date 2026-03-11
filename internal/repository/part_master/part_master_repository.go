package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)
const (
	TABLE_PART_MASTER 			= "part_masters"
)

type SqlPartMasterRepository struct {
	db *goqu.Database
}

func NewSqlPartMasterRepository(DB *goqu.Database) PartMasterRepository {
	return &SqlPartMasterRepository{
		db: DB,
	}
}

func (cr *SqlPartMasterRepository) GetAllPartMasters() ([]models.PartMaster, error) {
	var skills []models.PartMaster
	err := cr.db.From(TABLE_PART_MASTER).ScanStructs(&skills)
	if err != nil {
		return nil, err
	}
	
	return skills, nil
}

func (cr *SqlPartMasterRepository) CreatePartMaster(params v1dto.PartMasterParamsInput) error {
	_, err := cr.db.From(TABLE_PART_MASTER).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (cr *SqlPartMasterRepository) FindPartMasterById(id int) (skill models.PartMaster, err error) {

	found, err := cr.db.From(TABLE_PART_MASTER).Where(goqu.C("id").Eq(id)).ScanStruct(&skill)
	if err != nil {
		return models.PartMaster{}, err
	}

	if !found {
		return models.PartMaster{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found part master.")
	}

	return skill, nil
}

func (cr *SqlPartMasterRepository) UpdatePartMasterById(id int, params goqu.Record) error {

	_, err := cr.db.From(TABLE_PART_MASTER).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}