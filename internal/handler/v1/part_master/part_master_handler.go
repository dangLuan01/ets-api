package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/part_master"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type PartMasterHandler struct {
	service v1service.PartMasterService
}

type EditPartMasterParams struct {
	Id int `uri:"id" binding:"required"`
}

func NewPartMasterHandler(service v1service.PartMasterService) *PartMasterHandler {
	return &PartMasterHandler {
		service: service,
	}
}

func (ch *PartMasterHandler) GetAllPartMasters(ctx *gin.Context) {
	var params v1dto.GetAllPartMasterParams
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

	partMasters, totalRecords, err := ch.service.GetAllPartMasters(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, partMasters)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (ch *PartMasterHandler) CreatePartMaster(ctx *gin.Context) {
	var params v1dto.PartMasterParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreatePartMaster(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *PartMasterHandler) EditPartMaster(ctx *gin.Context) {
	var params EditPartMasterParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	skill, err := ch.service.EditPartMaster(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", skill)
}

func (ch *PartMasterHandler) UpdatePartMaster(ctx *gin.Context) {
	var params v1dto.PartMasterParamsUpdate
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdatePartMaster(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *PartMasterHandler) DeletePartMaster(ctx *gin.Context) {}