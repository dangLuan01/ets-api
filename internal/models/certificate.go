package models

type Certificate struct {
	Id 				int 					`json:"id" db:"id"`
	Code 			string 					`json:"code" db:"code"`
	Name 			string 					`json:"name" db:"name"`
	Description 	*string 				`json:"description" db:"description"`
	Status 			int 					`json:"status" db:"status"`
}