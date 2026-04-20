package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/post"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service v1service.PostService
}

type EditPostParams struct {
	Id int `uri:"id" binding:"required"`
}

type UpdatePostParams struct {
	Id 	int `uri:"id" binding:"required"`
	Post v1dto.PostParamsInput
}

func NewPostHandler(service v1service.PostService) *PostHandler {
	return &PostHandler {
		service: service,
	}
}

func (ch *PostHandler) GetAllPosts(ctx *gin.Context) {
	var params v1dto.GetAllPostParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 20
	}

	certificates, totalRecords, err := ch.service.GetAllPosts(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, certificates)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (ch *PostHandler) CreatePost(ctx *gin.Context) {
	var params v1dto.PostParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreatePost(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *PostHandler) EditPost(ctx *gin.Context) {
	var params EditPostParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	certificate, err := ch.service.EditPost(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", certificate)
}

func (ch *PostHandler) UpdatePost(ctx *gin.Context) {
	var params v1dto.PostParamsUpdate
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdatePost(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *PostHandler) DeletePost(ctx *gin.Context) {}