package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/menu"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

type menuService struct {
	repo repository.MenuRepository
}

func NewMenuService(repo repository.MenuRepository) MenuService {
	return &menuService{
		repo: repo,
	}
}

func (ms *menuService) GetMenuHeader(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error) {
	
	menu, total, err := ms.repo.FindMenuByType(params)
	menuTree := utils.BuildTree(menu)

	if err != nil {
		return nil, 0, err
	}

	return menuTree, total, nil
}

func (ms *menuService) GetAllMenu(params v1dto.GetAllMenuParams) ([]models.Menu, int64, error) {
	return ms.repo.GetAllMenu(params)
}

func (ms *menuService) CreateMenu(params v1dto.MenuInputParams) error {
	return ms.repo.CreateMenu(params)
}

func (ms *menuService) EditMenu(id int) (models.Menu, error) {
	return ms.repo.FindMenuById(id)
}

func (ms *menuService) UpdateMenu(params v1dto.MenuUpdateParams) error {
	updateData := goqu.Record{}
	
	updateData["name"] 			= params.Name
	updateData["priority"] 		= params.Priority
	updateData["type"] 			= params.Type
	updateData["status"] 		= params.Status
	
	if params.Slug != nil {
		updateData["slug"] 		= params.Slug
	}
	if params.ParentId != nil {
		updateData["parent_id"] = params.ParentId
	}
	
	return ms.repo.UpdateMenuById(params.Id, updateData)
}

func (ms *menuService) GetMenuStructure() ([]*v1dto.MenuDTO, error) {

	menuStructure, err := ms.repo.FindMenuStructure()
	if err != nil {
		return nil, err
	}
	
	return utils.BuildTree(menuStructure), nil
}