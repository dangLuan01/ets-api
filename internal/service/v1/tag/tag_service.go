package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/tag"
	"github.com/doug-martin/goqu/v9"
)

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{
		repo: repo,
	}
}

func (ts *tagService) GetAllTags(params v1dto.GetAllTagParams) ([]models.Tag, int64, error) {
	
	return ts.repo.GetAllTags(params)
}

func (ts *tagService) CreateTag(params v1dto.TagParamsInput) error {
	return ts.repo.CreateTag(params)
}

func (ts *tagService) EditTag(id int) (models.Tag, error) {
	return ts.repo.FindTagById(id)
}

func (ts *tagService) UpdateTag(params v1dto.TagParamsUpdate) error {
	updateData := goqu.Record{}
	updateData["name"] 		= params.Name
	updateData["slug"] 		= params.Slug
	updateData["status"] 	= params.Status
	
	return ts.repo.UpdateTagById(params.Id, updateData)
}