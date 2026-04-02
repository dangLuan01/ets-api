package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	repository "github.com/dangLuan01/ets-api/internal/repository/menu"
	"github.com/dangLuan01/ets-api/internal/utils"
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
	switch params.Type {
	case "header":
		headers, total, err := ms.repo.FindMenuHeader(params)
		headerTree := utils.BuildTree(headers)

		if err != nil {
			return nil, 0, err
		}

		return headerTree, total, nil
	default:
		footers, total, err := ms.repo.FindMenuFooter(params)
		footerTree := utils.BuildTree(footers)

		if err != nil {
			return nil, 0, err
		}

		return footerTree, total, nil
	}
}