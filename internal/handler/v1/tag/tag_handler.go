package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/tag"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	service v1service.TagService
}

type EditTagParams struct {
	Id int `uri:"id" binding:"required"`
}

type UpdateTagParams struct {
	Id 	int `uri:"id" binding:"required"`
	Tag v1dto.TagParamsInput
}

func NewTagHandler(service v1service.TagService) *TagHandler {
	return &TagHandler {
		service: service,
	}
}

func (ch *TagHandler) GetAllTags(ctx *gin.Context) {
	var params v1dto.GetAllTagParams
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

	certificates, totalRecords, err := ch.service.GetAllTags(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, certificates)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (ch *TagHandler) CreateTag(ctx *gin.Context) {
	var params v1dto.TagParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreateTag(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *TagHandler) EditTag(ctx *gin.Context) {
	var params EditTagParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	certificate, err := ch.service.EditTag(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", certificate)
}

func (ch *TagHandler) UpdateTag(ctx *gin.Context) {
	var params v1dto.TagParamsUpdate
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdateTag(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *TagHandler) DeleteTag(ctx *gin.Context) {}