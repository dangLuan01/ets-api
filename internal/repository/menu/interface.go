package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type MenuRepository interface {
	//====================CLIENT===================
	FindMenuByType(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error)
	//====================ADMIN====================
	GetAllMenu(params v1dto.GetAllMenuParams) ([]models.Menu, int64, error)
	CreateMenu(params v1dto.MenuInputParams) error
	FindMenuById(id int) (models.Menu, error)
	UpdateMenuById(id int, params goqu.Record) error
	FindMenuStructure() ([]*v1dto.MenuDTO, error)
}