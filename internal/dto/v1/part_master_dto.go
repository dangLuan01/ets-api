package v1dto

type PartMasterParamsInput struct {
	SkillId 	int 		`json:"skill_id" db:"skill_id" binding:"required"`
	Name 		string 		`json:"name" db:"name" binding:"required,max=100"`
	PartNumber 	int  		`json:"part_number" db:"part_number" binding:"required,maxInt=100"`
}

type PartMasterParamsUpdate struct {
	Id 			int 		`json:"id" binding:"required"`
	SkillId 	*int 		`json:"skill_id" binding:"omitempty"`
	Name 		*string 	`json:"name" binding:"omitempty"`
	PartNumber 	*int 		`json:"part_number" binding:"omitempty"`
	Status 		*int 		`json:"status" binding:"omitempty,oneof=0 1"`
}