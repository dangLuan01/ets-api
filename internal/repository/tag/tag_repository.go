package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_TAG = "post_tags"
)

type SqlTagRepository struct {
	db *goqu.Database
}

func NewSqlTagRepository(DB *goqu.Database) TagRepository {
	return &SqlTagRepository{
		db: DB,
	}
}

func (cr *SqlTagRepository) GetAllTags(params v1dto.GetAllTagParams) ([]models.Tag, int64, error) {
	var tags []models.Tag
	ds := cr.db.From(TABLE_TAG)
	if params.Name != "" {
		ds = ds.Where(goqu.C("name").ILike("%" + params.Name + "%"))
	}

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}
	
	if err := ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).ScanStructs(&tags); err != nil {
		return nil, 0, err
	}
	
	return tags, totalRecords, nil
}

func (cr *SqlTagRepository) CreateTag(params v1dto.TagParamsInput) error {
	_, err := cr.db.From(TABLE_TAG).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (cr *SqlTagRepository) FindTagById(id int) (tag models.Tag, err error) {

	found, err := cr.db.From(TABLE_TAG).Where(goqu.C("id").Eq(id)).ScanStruct(&tag)
	if err != nil {
		return models.Tag{}, err
	}

	if !found {
		return models.Tag{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found tag.")
	}

	return tag, nil
}

func (cr *SqlTagRepository) UpdateTagById(id int, params goqu.Record) error {

	_, err := cr.db.From(TABLE_TAG).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}