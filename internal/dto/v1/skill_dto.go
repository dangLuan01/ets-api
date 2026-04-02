package v1dto

type GetAllSkillParams struct {
	Page int32 `form:"page" binding:"omitempty,min=1"`
	Limit int32 `form:"limit" binding:"omitempty,min=1,max=50"`
}

type SkillParamsInput struct {
	CertId 		int 		`json:"cert_id" db:"cert_id" binding:"required"`
	Code 		string 		`json:"code" db:"code" binding:"required,max=100"`
	Name 		string  	`json:"name" db:"name" binding:"required,max=255"`
	OrderIndex 	int 		`json:"order_index" db:"order_index" binding:"required"`
}

type SkillParamsUpdate struct {
	Id 			int 		`json:"id" binding:"required"`
	CertId 		*int 		`json:"cert_id" binding:"omitempty"`
	Code 		*string 	`json:"code" binding:"omitempty"`
	Name 		*string 	`json:"name" binding:"omitempty"`
	OrderIndex 	*int 		`json:"order_index" binding:"omitempty"`
	Status 		*int 		`json:"status" binding:"omitempty,oneof=0 1"`
}