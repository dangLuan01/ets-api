package v1dto

type CertificateParamsInput struct {
	Code 		string  `json:"code" db:"code" binding:"required,max=100"`
	Name 		string  `json:"name" db:"name" binding:"required,max=255"`
	Description *string `json:"description" db:"description" binding:"omitempty,max=255"`
}

type CertificateParamsUpdate struct {
	Id 			int 	`json:"id" binding:"required"`
	Code 		*string `json:"code" binding:"omitempty"`
	Name 		*string `json:"name" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
	Status 		*int 	`json:"status" binding:"omitempty,oneof=0 1"`
}