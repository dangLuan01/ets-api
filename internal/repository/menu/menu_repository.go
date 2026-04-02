package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
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

func (mr *SqlMenuRepository) FindMenuHeader(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error) {
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

func (mr *SqlMenuRepository) FindMenuFooter(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error) {
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