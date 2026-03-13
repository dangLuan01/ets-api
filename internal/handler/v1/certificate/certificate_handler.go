package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/certificate"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type CertificateHandler struct {
	service v1service.CertificateService
}

type EditCertificateParams struct {
	Id int `uri:"id" binding:"required"`
}

type UpdateCertificateParams struct {
	Id int `uri:"id" binding:"required"`
	Certificate v1dto.CertificateParamsInput `json:""`
}

func NewCertificateHandler(service v1service.CertificateService) *CertificateHandler {
	return &CertificateHandler {
		service: service,
	}
}

func (ch *CertificateHandler) GetAllCertificates(ctx *gin.Context) {
	var params v1dto.GetAllCertificateParams
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

	certificates, totalRecords, err := ch.service.GetAllCertificates(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, certificates)

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (ch *CertificateHandler) CreateCertificate(ctx *gin.Context) {
	var params v1dto.CertificateParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreateCertificate(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *CertificateHandler) EditCertificate(ctx *gin.Context) {
	var params EditCertificateParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	certificate, err := ch.service.EditCertificate(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", certificate)
}

func (ch *CertificateHandler) UpdateCertificate(ctx *gin.Context) {
	var params v1dto.CertificateParamsUpdate
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdateCertificate(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *CertificateHandler) DeleteCertificate(ctx *gin.Context) {}