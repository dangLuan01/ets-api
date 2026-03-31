package v1service

import v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"

type MenuService interface {
	GetMenuHeader(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error)
}