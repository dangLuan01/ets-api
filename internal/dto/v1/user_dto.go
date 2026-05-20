package v1dto

import (
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/google/uuid"
)

type UserDTO struct {
	UserName   	string 		`json:"username"`
	Email  		string 		`json:"email"`
	Avatar		*string		`json:"avatar"`
	Target		int			`json:"target"`
	ExamDate	*string		`json:"exam_date"`
	HasPassword bool		`json:"has_password"`
}

type CreateUserInput struct {
	UUID   		uuid.UUID 	`json:"uuid"`
	Name     	string 		`json:"name" binding:"required"`
	Email    	string 		`json:"email" binding:"required,email"`
	Password 	string 		`json:"password" binding:"required,password,min=8"`
	Status   	int8   		`json:"status" binding:"required,oneof=1 2"`
	Role    	int8   		`json:"role" binding:"required,oneof=1 2"`
}

type UpdateUserInput struct {
	Name     	string 		`json:"name" binding:"required"`
	Target		int			`json:"target" binding:"required,minInt=1,maxInt=990"`
	ExamDate	*string		`json:"exam_date" binding:"omitempty"`
}

type UpdatePasswordInput struct {
	HasPassword 	bool 	`json:"has_password" binding:"omitempty"`
	CurrentPassword string 	`json:"current_password" binding:"omitempty"`
	NewPassword 	string 	`json:"new_password" binding:"required,password,min=8"`
	ConfirmPassword string 	`json:"confirm_password" binding:"required"`
}

type UserPayload struct {
	UserUUID 	uuid.UUID 	`json:"user_uuid"`
	Email 		string 		`json:"email"`
	Role 		int8 		`json:"role"`
}

type GetAttemptByUserUuuidParams struct {
	Page int32 `form:"page" binding:"omitempty,min=1"`
	Limit int32 `form:"limit" binding:"omitempty,min=1,max=50"`
	Status int	`form:"status" binding:"omitempty,oneof=1 2"`
}

func (input * CreateUserInput) MapCreateInputToModel() models.User {
	return models.User{
		UserName: input.Name,
		Email: input.Email,
		PasswordHash: &input.Password,
		Status: input.Status,
		Role: input.Role,
	}
}

func MapUserDTO(user models.User) *UserDTO {
	hasPassword := false

	if user.PasswordHash != nil {
		hasPassword = true	
	}

	return &UserDTO{
		UserName: user.UserName,
		Email: user.Email,
		Avatar: user.Avatar,
		Target: user.Target,
		ExamDate: user.ExamDate,
		HasPassword: hasPassword,
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