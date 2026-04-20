package v1dto

type GetAllTagParams struct {
	Page 		int32 	`form:"page" binding:"omitempty,min=1"`
	Limit 		int32 	`form:"limit" binding:"omitempty,min=1,max=50"`
	Name 		string 	`form:"name" binding:"omitempty,max=255"`
}

type TagParamsInput struct {
	Name 		string  `json:"name" db:"name" binding:"required,max=255"`
	Slug 		string  `json:"slug" db:"slug" binding:"required,max=255"`
	Status 		int  	`json:"status" db:"status" binding:"required,oneof=0 1"`
}

type TagParamsUpdate struct {
	Id 			int 	`json:"id" binding:"required"`
	Name 		string  `json:"name" binding:"required,max=255"`
	Slug 		string  `json:"slug" binding:"required,max=255"`
	Status 		int  	`json:"status" binding:"omitempty,oneof=0 1"`
}