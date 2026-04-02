package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_CATEGORY	= "categories"
)

type SqlCategoryRepository struct {
	db *goqu.Database
}

func NewSqlCategoryRepository(DB *goqu.Database) CategoryRepository {
	return &SqlCategoryRepository{
		db: DB,
	}
}

func (cr *SqlCategoryRepository) GetAllCategory(params v1dto.GetAllCategoryParams) ([]models.Category, int64, error) {
	var category []models.Category
	ds := cr.db.From(TABLE_CATEGORY)

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}

	if err:= ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).Order(goqu.C("created_at").Desc()).ScanStructs(&category); err != nil {
		return nil, 0, err
	}
	
	return category, totalRecords, nil
}

func (cr *SqlCategoryRepository) CreateCategory(params v1dto.CategoryInputParams) error {
	_, err := cr.db.From(TABLE_CATEGORY).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (cr *SqlCategoryRepository) FindCategoryById(id int) (category models.Category, err error) {

	found, err := cr.db.From(TABLE_CATEGORY).Where(goqu.C("id").Eq(id)).ScanStruct(&category)
	if err != nil {
		return models.Category{}, err
	}

	if !found {
		return models.Category{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found skill.")
	}

	return category, nil
}

func (cr *SqlCategoryRepository) UpdateCategoryById(id int, params goqu.Record) error {
	
	_, err := cr.db.From(TABLE_CATEGORY).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}

func (er *SqlCategoryRepository) FindCategoryStructure() ([]*v1dto.CategoryStructure, error) {
	var categoryStructure []*v1dto.CategoryStructure

	ds := er.db.From(goqu.T(TABLE_CATEGORY)).
		Order(goqu.C("priority").Asc())

	err := ds.ScanStructs(&categoryStructure)
	if err != nil {
		return nil, err
	}

	return categoryStructure, nil
}