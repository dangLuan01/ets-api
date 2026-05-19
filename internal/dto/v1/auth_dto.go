package v1dto

import (
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/google/uuid"
)

type LoginInput struct {
	Email    		string 	`json:"email" binding:"required,email"`
	Password 		string 	`json:"password" binding:"required,min=8"`
	Token			string	`json:"token" binding:"required"`
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
	UserName 		string 	`json:"username" binding:"required,max=50"`
	Email 	 		string 	`json:"email" binding:"required,email,max=50"`
	Password 		string 	`json:"password" binding:"required,min=8"`
	Target 			int 	`json:"target" binding:"required,min=10,max=990"`
	Token			string	`json:"token" binding:"required"`
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
		PasswordHash: &user.Password,
		Target: user.Target,
		Role: 2,
		Status: 1,
	}
}

func Oauth2DTOToModel(uuid uuid.UUID, provider string, user models.User) models.User {
	return models.User{
		UUID: uuid,
		UserName: user.UserName,
		Email: user.Email,
		Avatar: user.Avatar,
		Provider: &provider,
		OpenID: user.OpenID,
		Target: 990,
		Role: 2,
		Status: 1,
	}
}