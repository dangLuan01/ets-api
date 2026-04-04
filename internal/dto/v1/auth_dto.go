package v1dto

import (
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/google/uuid"
)

type LoginInput struct {
	Email    		string `json:"email" binding:"required,email"`
	Password 		string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	AccessToken 	string 	`json:"access_token"`
	RefreshToken 	string	`json:"refresh_token"`
	ExpiresIn 		int 	`json:"expires_in"`
}

type RefreshTokenInput struct {
	RefreshToken 	string `json:"refresh_token" binding:"required"`
}

type RegisterInput struct {
	UserName 		string `json:"username" binding:"required,max=50"`
	Email 	 		string `json:"email" binding:"required,email,max=50"`
	Password 		string `json:"password" binding:"required,min=8"`
}

type RequestOTPInput struct {
	Code string `json:"code" binding:"required,max=6"`
}

type EncryptedPayload struct {
	UserUUID 	uuid.UUID 	`json:"user_uuid"`
	Email 		string 		`json:"email"`
	Role 		int8 		`json:"role"`
}

func RegisterDTOToModel(uuid uuid.UUID, user RegisterInput) models.User {
	return models.User{
		UUID: uuid,
		UserName: user.UserName,
		Email: user.Email,
		PasswordHash: user.Password,
		Role: 2,
		Status: 1,
	}
}