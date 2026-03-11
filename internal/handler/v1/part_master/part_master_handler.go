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
	skills, err := ch.service.GetAllPartMasters()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", skills)
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