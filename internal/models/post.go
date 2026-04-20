package models

type Post struct {
	Id 				int 	`json:"id" db:"id"`
	Name 			string  `json:"name" db:"name"`
	Slug 			string  `json:"slug" db:"slug"`
	Content 		string  `json:"content" db:"content"`
	Summary 		string  `json:"summary" db:"summary"`
	ThumbnailUrl 	string  `json:"thumbnail_url" db:"thumbnail_url"`
	Status 			int  	`json:"status" db:"status"`
	Tags			[]int	`json:"tags" db:"-"`
}