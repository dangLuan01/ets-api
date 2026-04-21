package repository

import (
	"encoding/json"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_POST 		= "posts"
	TABLE_TAG 		= "post_tags"
	TABLE_POST_TAG 	= "post_tag_mappings"
)

type SqlPostRepository struct {
	db *goqu.Database
}

func NewSqlPostRepository(DB *goqu.Database) PostRepository {
	return &SqlPostRepository{
		db: DB,
	}
}

func (cr *SqlPostRepository) GetAllPosts(params v1dto.GetAllPostParams) ([]models.Post, int64, error) {
	var posts []models.Post
	ds := cr.db.From(TABLE_POST)
	if params.Name != "" {
		ds = ds.Where(goqu.C("name").ILike("%" + params.Name + "%"))
	}

	totalRecords, err := ds.Count()
	if err != nil {
		return nil, 0, err
	}
	
	if err := ds.Offset((uint(params.Page) - 1) * uint(params.Limit)).Limit(uint(params.Limit)).ScanStructs(&posts); err != nil {
		return nil, 0, err
	}
	
	return posts, totalRecords, nil
}

func (cr *SqlPostRepository) CreatePost(tx *goqu.TxDatabase, params v1dto.PostParamsInput) (int64, error) {
	resp, err := tx.From(TABLE_POST).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return 0, err
	}

	postId, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (cr *SqlPostRepository) FindPostById(id int) (post models.Post, err error) {

	found, err := cr.db.From(TABLE_POST).Where(goqu.C("id").Eq(id)).ScanStruct(&post)
	if err != nil {
		return models.Post{}, err
	}

	if !found {
		return models.Post{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found post.")
	}

	var tags []int
	err = cr.db.From(TABLE_POST_TAG).
		Select(
			goqu.C("tag_id"),
		).
		Where(goqu.C("post_id").Eq(id)).
		ScanVals(&tags)
	if err != nil {
		return models.Post{}, err
	}

	post.Tags = tags

	return post, nil
}

func (cr *SqlPostRepository) UpdatePostById(tx *goqu.TxDatabase, id int, params goqu.Record) error {

	_, err := tx.From(TABLE_POST).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}

func (cr *SqlPostRepository) InsertPostTag(tx *goqu.TxDatabase, rows []goqu.Record) error {
	_, err := tx.Insert(TABLE_POST_TAG).Rows(rows).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (cr *SqlPostRepository) DeletePostTag(tx *goqu.TxDatabase, postId int) error {
	_, err := tx.Delete(TABLE_POST_TAG).
	Where(
		goqu.C("post_id").Eq(postId),
	).Executor().Exec()
	
	if err != nil {
		return err
	}

	return nil
}
//=======================================Client===============================
func (cr *SqlPostRepository) FindAllPosts(params v1dto.GetAllPostParams) ([]v1dto.PostDTO, int64, error) {
	var posts []v1dto.PostDTO
	countDs := cr.db.From(goqu.T(TABLE_POST).As("p"))

	totalRecords, err := countDs.Count()
	if err != nil {
		return nil, 0, err
	}
	
	postSub := cr.db.From(goqu.T(TABLE_POST).As("p")).
	Select(
		goqu.I("p.id"),
		goqu.I("p.name"),
		goqu.I("p.slug"),
		goqu.I("p.summary"),
		goqu.I("p.thumbnail_url"),
		goqu.I("p.view_count"),
		goqu.I("p.updated_at"),
	).
	Order(goqu.I("p.updated_at").Desc()).
	Offset((uint(params.Page) - 1) * uint(params.Limit)).
	Limit(uint(params.Limit))

	tagSub := cr.db.From(goqu.T(TABLE_POST_TAG).As("pt")).
	Select(
		goqu.I("pt.post_id"),
		goqu.L(`
			JSON_ARRAYAGG(
				JSON_OBJECT(
					'name', t.name,
					'slug', t.slug
				)
			)
		`).As("tags"),
	).
	Join(goqu.T(TABLE_TAG).As("t"),
		goqu.On(goqu.I("t.id").Eq(goqu.I("pt.tag_id"))),
	).
	GroupBy(goqu.I("pt.post_id"))

	ds := cr.db.From(postSub.As("p")).
	Select(
		goqu.I("p.name"),
		goqu.I("p.slug"),
		goqu.I("p.summary"),
		goqu.I("p.thumbnail_url"),
		goqu.I("p.view_count"),
		goqu.I("p.updated_at"),
		goqu.COALESCE(goqu.I("tag_data.tags"), goqu.L("JSON_ARRAY()")).As("tags"),
	).
	LeftJoin(tagSub.As("tag_data"),
		goqu.On(goqu.I("tag_data.post_id").Eq(goqu.I("p.id"))),
	)

	if err := ds.ScanStructs(&posts); err != nil {
		return nil, 0, err
	}

	for i := range posts {
		if len(posts[i].TagsRaw) > 0 {
			_ = json.Unmarshal(posts[i].TagsRaw, &posts[i].Tags)
		}
	}

	return posts, totalRecords, nil
}

func (cr *SqlPostRepository) FindPostBySlug(slug string) (v1dto.PostDetailDTO, error) {
	var posts v1dto.PostDetailDTO
	
	postSub := cr.db.From(goqu.T(TABLE_POST).As("p")).
	Select(
		goqu.I("p.id"),
		goqu.I("p.name"),
		goqu.I("p.slug"),
		goqu.I("p.content"),
		goqu.I("p.summary"),
		goqu.I("p.thumbnail_url"),
		goqu.I("p.view_count"),
		goqu.I("p.updated_at"),
	).Where(goqu.I("p.slug").Eq(slug))

	tagSub := cr.db.From(goqu.T(TABLE_POST_TAG).As("pt")).
	Select(
		goqu.I("pt.post_id"),
		goqu.L(`
			JSON_ARRAYAGG(
				JSON_OBJECT(
					'name', t.name,
					'slug', t.slug
				)
			)
		`).As("tags"),
	).
	Join(goqu.T(TABLE_TAG).As("t"),
		goqu.On(goqu.I("t.id").Eq(goqu.I("pt.tag_id"))),
	).
	GroupBy(goqu.I("pt.post_id"))

	ds := cr.db.From(postSub.As("p")).
	Select(
		goqu.I("p.name"),
		goqu.I("p.slug"),
		goqu.I("p.content"),
		goqu.I("p.summary"),
		goqu.I("p.thumbnail_url"),
		goqu.I("p.view_count"),
		goqu.I("p.updated_at"),
		goqu.COALESCE(goqu.I("tag_data.tags"), goqu.L("JSON_ARRAY()")).As("tags"),
	).
	LeftJoin(tagSub.As("tag_data"),
		goqu.On(goqu.I("tag_data.post_id").Eq(goqu.I("p.id"))),
	)

	if _, err := ds.ScanStruct(&posts); err != nil {
		return v1dto.PostDetailDTO{}, err
	}

	if len(posts.TagsRaw) > 0 {
		_ = json.Unmarshal(posts.TagsRaw, &posts.Tags)
	}

	return posts, nil
}