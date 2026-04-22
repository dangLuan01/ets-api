package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type PostRepository interface {
	GetAllPosts(params v1dto.GetAllPostParams) ([]models.Post, int64, error)
	CreatePost(tx *goqu.TxDatabase, params v1dto.PostParamsInput) (int64, error)
	FindPostById(id int) (models.Post, error)
	UpdatePostById(tx *goqu.TxDatabase, id int, params goqu.Record) error
	InsertPostTag(tx *goqu.TxDatabase, rows []goqu.Record) error
	DeletePostTag(tx *goqu.TxDatabase, postId int) error
	//=========================Client=======================
	FindAllPosts(params v1dto.GetAllPostParams) ([]v1dto.PostDTO, int64, error)
	FindPostBySlug(slug string) (v1dto.PostDetailDTO, error)
	FindPostByTagSlug(slug string, page, limit int32) ([]v1dto.PostDTO, int64, error)
}