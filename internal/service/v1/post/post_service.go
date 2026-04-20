package v1service

import (
	"time"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/post"
	"github.com/doug-martin/goqu/v9"
)

type postService struct {
	db *goqu.Database
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository, DB *goqu.Database) PostService {
	return &postService{
		db: DB,
		repo: repo,
	}
}

func (ps *postService) GetAllPosts(params v1dto.GetAllPostParams) ([]models.Post, int64, error) {
	
	return ps.repo.GetAllPosts(params)
}

func (ps *postService) CreatePost(params v1dto.PostParamsInput) error {
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	postId, err := ps.repo.CreatePost(tx, params)
	if err != nil {
		return  err
	}

	if len(params.Tags) > 0 {
		rows := make([]goqu.Record, 0, len(params.Tags))
		for _, tagId := range params.Tags {
			rows = append(rows, goqu.Record{
				"post_id": postId,
				"tag_id": tagId,
			})
		}
		if err := ps.repo.InsertPostTag(tx, rows); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (ps *postService) EditPost(id int) (models.Post, error) {
	return ps.repo.FindPostById(id)
}

func (ps *postService) UpdatePost(params v1dto.PostParamsUpdate) error {
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	updateData := goqu.Record{}
	updateData["name"] 			= params.Name
	updateData["slug"] 			= params.Slug
	updateData["content"] 		= params.Content
	updateData["summary"] 		= params.Summary
	updateData["thumbnail_url"]	= params.ThumbnailUrl
	updateData["updated_at"] 	= time.Now()
	updateData["status"] 		= params.Status

	if err = ps.repo.UpdatePostById(tx, params.Id, updateData); err != nil {
		return err
	}

	if err = ps.repo.DeletePostTag(tx, params.Id); err != nil {
		return err
	}

	if len(params.Tags) > 0 {
		rows := make([]goqu.Record, 0, len(params.Tags))
		for _, tagId := range params.Tags {
			rows = append(rows, goqu.Record{
				"post_id": params.Id,
				"tag_id": tagId,
			})
		}
		if err := ps.repo.InsertPostTag(tx, rows); err != nil {
			return err
		}
	}

	return tx.Commit()
}