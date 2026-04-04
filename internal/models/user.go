package models

import (
	"github.com/google/uuid"
)

type User struct {
	UUID      		uuid.UUID 	`db:"uuid"`
	UserName  		string 		`db:"username"`
	Email     		string 		`db:"email"`
	PasswordHash  	string 		`db:"password_hash"`
	Role     		int8   		`db:"role"`
	Status    		int8   		`db:"status"`
}