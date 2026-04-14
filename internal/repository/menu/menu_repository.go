package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_MENU	= "menus"
)

type SqlMenuRepository struct {
	db *goqu.Database
}

func NewSqlMenuRepository(DB *goqu.Database) MenuRepository {
	return &SqlMenuRepository{
		db: DB,
	}
}

func (mr *SqlMenuRepository) FindMenuByType(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error) {
	var menus []*v1dto.MenuDTO
	ds := mr.db.From(TABLE_MENU).
		Where(goqu.C("type").Eq(params.Type))

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}

	ds = ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).Order(goqu.C("priority").Asc())

	if err := ds.ScanStructs(&menus); err != nil {
		return nil, 0, err
	}
	
	return menus, totalRecords, nil
}


func (mr *SqlMenuRepository) GetAllMenu(params v1dto.GetAllMenuParams) ([]models.Menu, int64, error) {
	var category []models.Menu
	ds := mr.db.From(TABLE_MENU)

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}

	if err:= ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).Order(goqu.C("priority").Desc()).ScanStructs(&category); err != nil {
		return nil, 0, err
	}
	
	return category, totalRecords, nil
}

func (mr *SqlMenuRepository) CreateMenu(params v1dto.MenuInputParams) error {
	_, err := mr.db.From(TABLE_MENU).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (mr *SqlMenuRepository) FindMenuById(id int) (category models.Menu, err error) {

	found, err := mr.db.From(TABLE_MENU).Where(goqu.C("id").Eq(id)).ScanStruct(&category)
	if err != nil {
		return models.Menu{}, err
	}

	if !found {
		return models.Menu{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found skill.")
	}

	return category, nil
}

func (mr *SqlMenuRepository) UpdateMenuById(id int, params goqu.Record) error {
	
	_, err := mr.db.From(TABLE_MENU).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}

func (mr *SqlMenuRepository) FindMenuStructure() ([]*v1dto.MenuDTO, error) {
	var categoryStructure []*v1dto.MenuDTO

	ds := mr.db.From(goqu.T(TABLE_MENU)).
		Order(goqu.C("priority").Asc())

	err := ds.ScanStructs(&categoryStructure)
	if err != nil {
		return nil, err
	}

	return categoryStructure, nil
}