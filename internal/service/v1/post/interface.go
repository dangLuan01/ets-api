package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type PostService interface {
	GetAllPosts(params v1dto.GetAllPostParams) ([]models.Post, int64, error)
	CreatePost(params v1dto.PostParamsInput) error
	EditPost(id int) (models.Post, error)
	UpdatePost(params v1dto.PostParamsUpdate) error
}