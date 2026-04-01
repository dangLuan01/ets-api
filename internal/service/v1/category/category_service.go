package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/category"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (cs *categoryService) GetAllCategory(params v1dto.GetAllCategoryParams) ([]models.Category, int64, error) {
	return cs.repo.GetAllCategory(params)
}

func (cs *categoryService) CreateCategory(params v1dto.CategoryInputParams) error {
	return cs.repo.CreateCategory(params)
}

func (cs *categoryService) EditCategory(id int) (models.Category, error) {
	return cs.repo.FindCategoryById(id)
}

func (cs *categoryService) UpdateCategory(params v1dto.CategoryUpdateParams) error {
	updateData := goqu.Record{}
	
	updateData["name"] 			= params.Name
	updateData["priority"] 		= params.Priority
	updateData["type"] 			= params.Type
	updateData["status"] 		= params.Status
	updateData["is_filterable"] = params.IsFilterable
	
	if params.Slug != nil {
		updateData["slug"] 		= params.Slug
	}
	if params.ParentId != nil {
		updateData["parent_id"] = params.Name
	}
	
	return cs.repo.UpdateCategoryById(params.Id, updateData)
}

func (es *categoryService) GetCategoryStructure() ([]*v1dto.CategoryStructure, error) {

	categoryStructure, err := es.repo.FindCategoryStructure()
	if err != nil {
		return nil, err
	}
	
	return utils.BuildTree(categoryStructure), nil
}