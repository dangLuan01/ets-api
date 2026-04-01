package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/category"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service v1service.CategoryService
}

type EditCategoryParams struct {
	Id int `uri:"id" binding:"required"`
}

func NewCategoryHandler(service v1service.CategoryService) *CategoryHandler {
	return &CategoryHandler {
		service: service,
	}
}

func (ch *CategoryHandler) GetAllCategory(ctx *gin.Context) {
	var params v1dto.GetAllCategoryParams
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

	categories, totalRecords, err := ch.service.GetAllCategory(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, categories)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (ch *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var params v1dto.CategoryInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreateCategory(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *CategoryHandler) EditCategory(ctx *gin.Context) {
	var params EditCategoryParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	skill, err := ch.service.EditCategory(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", skill)
}

func (ch *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var params v1dto.CategoryUpdateParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdateCategory(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *CategoryHandler) GetCategoryStructure(ctx *gin.Context) {
	categoryStructure, err := ch.service.GetCategoryStructure()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", categoryStructure)
}