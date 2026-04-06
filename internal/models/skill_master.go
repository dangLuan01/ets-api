package models

type SkillMaster struct {
	Id 				int 					`json:"id" db:"id"`
	CertId 			int 					`json:"cert_id" db:"cert_id"`
	Code 			string 					`json:"code" db:"code"`
	Name 			string 					`json:"name" db:"name"`
	OrderIndex      int     				`json:"order_index" db:"order_index"`
	Status 			int 					`json:"status" db:"status"`
}