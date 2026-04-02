package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      		uuid.UUID 	`db:"uuid"`
	UserName  		string 		`db:"username"`
	Email     		string 		`db:"email"`
	Password  		string 		`db:"password"`
	Is_Member 		int			`db:"is_member"`
	UploadCount 	int			`db:"upload_count"`
	Expried_date 	*time.Time	`db:"expried_date"`
	Level     		int8   		`db:"level"`
	Status    		int8   		`db:"status"`
}