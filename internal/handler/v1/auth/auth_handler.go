package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/auth"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService v1service.AuthService
}

func NewAuthHandler(service v1service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (ah *AuthHandler) Login(ctx *gin.Context) {

	var input v1dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	accessToken, refreshToken, expiresIn, err := ah.authService.Login(ctx, input.Email, input.Password)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	response := v1dto.LoginResponse{
		AccessToken: 	accessToken,
		RefreshToken: 	refreshToken,
		ExpiresIn: 		expiresIn,
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Login successfully!", response)
}

func (ah *AuthHandler) Logout(ctx *gin.Context) {
	var input v1dto.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ah.authService.Logout(ctx, input.RefreshToken); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusOK)

}

func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	var input v1dto.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	accessToken, refreshToken, expiresIn, err := ah.authService.RefreshToken(ctx, input.RefreshToken)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	
	response := v1dto.LoginResponse{
		AccessToken: 	accessToken,
		RefreshToken: 	refreshToken,
		ExpiresIn: 		expiresIn,
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Refresh token successfully!", response)
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
	var input v1dto.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ah.authService.Register(ctx, input); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusOK)
}