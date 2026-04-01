package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type CategoryService interface {
	GetAllCategory(params v1dto.GetAllCategoryParams) ([]models.Category, int64, error)
	CreateCategory(params v1dto.CategoryInputParams) error
	EditCategory(id int) (models.Category, error)
	UpdateCategory(params v1dto.CategoryUpdateParams) error
	GetCategoryStructure() ([]*v1dto.CategoryStructure, error)
}