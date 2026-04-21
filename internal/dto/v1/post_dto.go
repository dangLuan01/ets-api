package v1dto

type GetAllPostParams struct {
	Page 			int32 			`form:"page" binding:"omitempty,min=1"`
	Limit 			int32 			`form:"limit" binding:"omitempty,min=1,max=50"`
	Name 			string 			`form:"name" binding:"omitempty,max=255"`
}

type PostParamsInput struct {
	Name 			string  		`json:"name" db:"name" binding:"required,max=255"`
	Slug 			string  		`json:"slug" db:"slug" binding:"required,max=255"`
	Content 		string  		`json:"content" db:"content" binding:"required"`
	Summary 		string  		`json:"summary" db:"summary" binding:"required"`
	ThumbnailUrl 	string  		`json:"thumbnail_url" db:"thumbnail_url" binding:"required"`
	Status 			int  			`json:"status" db:"status" binding:"required,oneof=0 1"`
	Tags 			[]int			`json:"tags" db:"-" binding:"required"`
}

type PostParamsUpdate struct {
	Id 				int 			`json:"id" binding:"required"`
	Name 			string  		`json:"name" binding:"required,max=255"`
	Slug 			string  		`json:"slug" binding:"required,max=255"`
	Content 		string  		`json:"content" binding:"required"`
	Summary 		string  		`json:"summary" binding:"required"`
	ThumbnailUrl 	string  		`json:"thumbnail_url" binding:"required"`
	Status 			int  			`json:"status" binding:"omitempty,oneof=0 1"`
	Tags 			[]int			`json:"tags" db:"-" binding:"required"`
}

type PostDTO struct {
	Name 			string  		`json:"name" db:"name"`
	Slug 			string  		`json:"slug" db:"slug"`
	Summary 		string  		`json:"summary" db:"summary"`
	ThumbnailUrl 	string  		`json:"thumbnail_url" db:"thumbnail_url"`
	ViewCount		int				`json:"view_count" db:"view_count"`
	UpdatedAt		string			`json:"updated_at" db:"updated_at"`
	TagsRaw 		[]byte   		`json:"-" db:"tags"`
	Tags    		[]TagDTO 		`json:"tags" db:"-"`
}

type PostDetailDTO struct {
	Name 			string  		`json:"name" db:"name"`
	Slug 			string  		`json:"slug" db:"slug"`
	Content 		string  		`json:"content" db:"content"`
	Summary 		string  		`json:"summary" db:"summary"`
	ThumbnailUrl 	string  		`json:"thumbnail_url" db:"thumbnail_url"`
	ViewCount		int				`json:"view_count" db:"view_count"`
	UpdatedAt		string			`json:"updated_at" db:"updated_at"`
	TagsRaw 		[]byte   		`json:"-" db:"tags"`
	Tags    		[]TagDTO 		`json:"tags" db:"-"`
}

type TagDTO struct {
	Name 		string  			`json:"name" db:"name"`
	Slug 		string  			`json:"slug" db:"slug"`
}