package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type CategoryRepository interface {
	GetAllCategory(params v1dto.GetAllCategoryParams) ([]models.Category, int64, error)
	CreateCategory(params v1dto.CategoryInputParams) error
	FindCategoryById(id int) (models.Category, error)
	UpdateCategoryById(id int, params goqu.Record) error
	FindCategoryStructure() ([]*v1dto.CategoryStructure, error)
}