package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
)

type MenuRepository interface {
	FindMenuHeader(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error)
	FindMenuFooter(params v1dto.GetMenuParams) ([]*v1dto.MenuDTO, int64, error)
}