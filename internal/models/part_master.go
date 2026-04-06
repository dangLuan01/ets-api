package models

type PartMaster struct {
	Id 				int 					`json:"id" db:"id"`
	SkillId 		int 					`json:"skill_id" db:"skill_id"`
	PartNumber 		int 					`json:"part_number" db:"part_number"`
	Name 			string 					`json:"name" db:"name"`
	Status 			int 					`json:"status" db:"status"`
}