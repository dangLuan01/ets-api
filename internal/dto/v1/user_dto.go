package v1dto

import (
	"time"

	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/google/uuid"
)

type UserDTO struct {
	UserName   		string 		`json:"user_name"`
	Email  			string 		`json:"email"`
	Role  			string 		`json:"role"`
	Status 			string 		`json:"status"`
	Expried_date 	*time.Time	`json:"expried_date"`
	Is_Member		int			`json:"is_member"`
}

type CreateUserInput struct {
	UUID   		uuid.UUID 	`json:"uuid"`
	Name     	string 		`json:"name" binding:"required"`
	Email    	string 		`json:"email" binding:"required,email"`
	Password 	string 		`json:"password" binding:"required,min=8"`
	Status   	int8   		`json:"status" binding:"required,oneof=1 2"`
	Role    	int8   		`json:"role" binding:"required,oneof=1 2"`
}

type UpdateUserInput struct {
	UUID   uuid.UUID 	`json:"uuid"`
	Name     string 	`json:"name" binding:"required"`
	Email    string 	`json:"email" binding:"required,email"`
	Password string 	`json:"password" binding:"omitempty,min=8"`
	Status   int8   	`json:"status" binding:"omitempty,oneof=1 2"`
	Role    int8   	`json:"role" binding:"omitempty,oneof=1 2"`
}

type ChangerPasswordParams struct {
	Password string `json:"password" binding:"required,min=8"`
}

type UserPayload struct {
	UserUUID 	uuid.UUID 	`json:"user_uuid"`
	Email 		string 		`json:"email"`
	Role 		int8 		`json:"role"`
}

func (input * CreateUserInput) MapCreateInputToModel() models.User {
	return models.User{
		UserName: input.Name,
		Email: input.Email,
		PasswordHash: input.Password,
		Status: input.Status,
		Role: input.Role,
	}
}

func (input * UpdateUserInput) MapUpdateInputToModel() models.User {
	return models.User{
		UserName: input.Name,
		Email: input.Email,
		PasswordHash: input.Password,
		Status: input.Status,
		Role: input.Role,
	}
}

func MapUserDTO(user models.User) *UserDTO {
	return &UserDTO{
		UserName: user.UserName,
		Email: user.Email,
		Role: formatLevel(user.Role),
		Status: formatStatus(user.Status),
	}
}

func MapUsersDTO(users []models.User) []UserDTO {
	dtos := make([]UserDTO, 0, len(users))
	for _, user := range users {
		dtos = append(dtos, *MapUserDTO(user))
	}
	return dtos
}

func formatLevel(role int8) string {
	switch role {
	case 1:
		return "Admin"
	default :
		return "Customer"
	}
}

func formatStatus(status int8) string {
	switch status {
	case 1:
		return "Active"
	default :
		return "Hidden"
	}
}