package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/menu"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	service v1service.MenuService
}

type EditMenuParams struct {
	Id int `uri:"id" binding:"required"`
}

func NewMenuHandler(service v1service.MenuService) *MenuHandler {
	return &MenuHandler {
		service: service,
	}
}

func (mh *MenuHandler) GetMenuHeader(ctx *gin.Context) {
	var params v1dto.GetMenuParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if params.Limit <= 0 {
		params.Limit = 10
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	menus, totalRecords, err := mh.service.GetMenuHeader(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, menus)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

//===========================================ADMIN=============================================

func (ch *MenuHandler) GetAllMenu(ctx *gin.Context) {
	var params v1dto.GetAllMenuParams
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

	categories, totalRecords, err := ch.service.GetAllMenu(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, categories)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (ch *MenuHandler) CreateMenu(ctx *gin.Context) {
	var params v1dto.MenuInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreateMenu(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *MenuHandler) EditMenu(ctx *gin.Context) {
	var params EditMenuParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	skill, err := ch.service.EditMenu(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", skill)
}

func (ch *MenuHandler) UpdateMenu(ctx *gin.Context) {
	var params v1dto.MenuUpdateParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdateMenu(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *MenuHandler) GetMenuStructure(ctx *gin.Context) {
	menuStructure, err := ch.service.GetMenuStructure()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", menuStructure)
}