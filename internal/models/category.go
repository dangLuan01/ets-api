package models

type Category struct {
	Id 				int 					`json:"id" db:"id"`
	ParentId 		*int 					`json:"parent_id" db:"parent_id"`
	Name			string 					`json:"name" db:"name"`
	Slug			*string					`json:"slug" db:"slug"`
	Type			string					`json:"type" db:"type"`
	Status 			int						`json:"status" db:"status"`
	IsFilterable 	int 					`json:"is_filterable" db:"is_filterable"`
	Priority 		int 					`json:"priority" db:"priority"`
}