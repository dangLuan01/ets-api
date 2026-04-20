package models

type Tag struct {
	Id 			int 	`json:"id" db:"id"`
	Name 		string  `json:"name" db:"name"`
	Slug 		string  `json:"slug" db:"slug"`
	Status 		int  	`json:"status" db:"status"`
}