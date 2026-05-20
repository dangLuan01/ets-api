package models

import (
	"github.com/google/uuid"
)

type User struct {
	UUID      		uuid.UUID 	`db:"uuid"`
	UserName  		string 		`db:"username"`
	Email     		string 		`db:"email"`
	PasswordHash  	*string 	`db:"password_hash"`
	Avatar			*string		`db:"avatar"`
	Provider		string		`db:"provider"`
	OpenID			*string		`db:"openid"`
	Target     		int 		`db:"target"`
	Role     		int8   		`db:"role"`
	ExamDate		*string		`db:"exam_date"`
	Status    		int8   		`db:"status"`
}

type UserUpdate struct {
	UserName  		string 		`db:"username"`
	Target     		int 		`db:"target"`
	ExamDate		*string		`db:"exam_date"`
}

type GoogleUser struct {
	Sub 			string 		`json:"sub"`
	Name 			string 		`json:"name"`
	Picture			string 		`json:"picture"`
	Email 			string 		`json:"email"`
}

type FacebookUser struct {
	Id 				string 		`json:"id"`
	Name 			string 		`json:"name"`	
	Email 			string 		`json:"email"`
	Picture			struct {
		Data struct {
			Url string `json:"url"`
		} `json:"data"`
	} 	`json:"picture"`
}