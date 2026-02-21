package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1"
	"github.com/google/uuid"

	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service v1service.UserService
}

type GetUserByUUIDParam struct{
	Uuid uuid.UUID `uri:"uuid" binding:"uuid"`
}

type GetUserByUUID struct{
	Uuid string `uri:"uuid" binding:"uuid"`
}

type CheckStatusParams struct {
	Uuid string `form:"uuid" binding:"required,uuid"`
}

func NewUserHandler(service v1service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
func (uh *UserHandler) GetAllUser(ctx *gin.Context)  {
	users, err := uh.service.GetAllUser()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully" ,v1dto.MapUsersDTO(users))
	
}
func (uh *UserHandler) GetUserByUUID(ctx *gin.Context)  {
	user, err := uh.service.GetUserByUUID(ctx)

	if err != nil {

		utils.ResponseError(ctx, err)
		return
	}
	
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully", v1dto.MapUserDTO(user))
}
func (uh *UserHandler) CreateUser(ctx *gin.Context) {

	var input v1dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))

		return
	}
	user := input.MapCreateInputToModel()
	
	createUser, err := uh.service.CreateUser(user)
	if err != nil {

		utils.ResponseError(ctx, err)

		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "Successfully", v1dto.MapUserDTO(createUser))
}
func (uh *UserHandler) UpdateUser(ctx *gin.Context)  {
	var param GetUserByUUIDParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return 
	}

	var input v1dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))

		return
	}

	user := input.MapUpdateInputToModel()

	updateUser, err := uh.service.UpdateUser(param.Uuid, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK ,"Successfully", updateUser)
}
func (uh *UserHandler) DeleteUser(ctx *gin.Context)  {
	var param GetUserByUUIDParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return 
	}
	if err := uh.service.DeleteUser(param.Uuid); err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	
	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (uh *UserHandler) CheckStatus(ctx *gin.Context) {
	var param CheckStatusParams
	if err := ctx.ShouldBindBodyWithJSON(&param); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := uh.service.CheckStatus(ctx, param.Uuid); err != nil {
		utils.ResponseError(ctx , err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully!")
}

func (uh *UserHandler) ChangePassword(ctx *gin.Context) {
	var params v1dto.ChangerPasswordParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := uh.service.ChangePassword(ctx, params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (uh *UserHandler) UpdateCountUpload(ctx *gin.Context)  {
	var params GetUserByUUID
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := uh.service.UpdateCountUpload(params.Uuid); err != nil {
		utils.ResponseError(ctx ,err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}