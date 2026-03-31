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