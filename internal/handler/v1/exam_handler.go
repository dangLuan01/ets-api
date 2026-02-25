package v1handler

import (
	"net/http"

	v1service "github.com/dangLuan01/ets-api/internal/service/v1"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	service v1service.ExamService
}

type GetIdExamParams struct {
	Id string `uri:"id" binding:"required"`
}

func NewExamHandler(service v1service.ExamService) *ExamHandler {
	return &ExamHandler {
		service: service,
	}
}

func (rh *ExamHandler) FindExamById(ctx *gin.Context) {
	var params GetIdExamParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	exam, err := rh.service.FindExamById(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", exam)
}