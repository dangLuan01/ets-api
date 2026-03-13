package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/question"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	service v1service.QuestionService
}

func NewQuestionHandler(service v1service.QuestionService) *QuestionHandler {
	return &QuestionHandler {
		service: service,
	}
}

func (qh *QuestionHandler) CreateQuestion(ctx *gin.Context) {
	var params v1dto.QuestionParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := qh.service.CreateQuestion(params);err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (qh *QuestionHandler) CreateQuestionGroup(ctx *gin.Context) {
	var params v1dto.QuestionGroupParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}
	
	if err := qh.service.CreateQuestionGroup(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}