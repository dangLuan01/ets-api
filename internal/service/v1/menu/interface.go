package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type MenuService interface {
	//==============CLIENT=================
	GetMenuHeader(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error)
	//==============ADMIN==================
	GetAllMenu(params v1dto.GetAllMenuParams) ([]models.Menu, int64, error)
	CreateMenu(params v1dto.MenuInputParams) error
	EditMenu(id int) (models.Menu, error)
	UpdateMenu(params v1dto.MenuUpdateParams) error
	GetMenuStructure() ([]*v1dto.MenuDTO, error)
}