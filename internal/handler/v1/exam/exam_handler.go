package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/exam"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	service v1service.ExamService
}

type GetIdExamParams struct {
	Id int `uri:"id" binding:"required"`
}

type GetExamPartParams struct {
	ExamId int `uri:"id" binding:"required"`
	PartId int `uri:"part_id" binding:"required"`
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

func (rh *ExamHandler) CalculateScoreExam(ctx *gin.Context) {
	var params v1dto.QuestionAnswerInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	data, err := rh.service.CalculateScoreExam(params)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", v1dto.MapDetailExamScoreDTO(data))
}

// --- HANDLER CHO ADMIN (CRUD EXAM) ---

func (rh *ExamHandler) GetAllExams(ctx *gin.Context){
	var params v1dto.GetAllExamParams
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

	exams, totalRecords, err := rh.service.GetAllExams(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, exams)
	
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (rh *ExamHandler) CreateExam(ctx *gin.Context){
	var params v1dto.CreateExamInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := rh.service.CreateExam(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (rh *ExamHandler) EditExam(ctx *gin.Context){
	var params GetIdExamParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	exam, err := rh.service.EditExamById(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", exam)
}

func (rh *ExamHandler) UpdateExam(ctx *gin.Context){
	var params v1dto.UpdateExamInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}
	
	if err := rh.service.UpdateExam(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (rh *ExamHandler) CreatePartDirection(ctx *gin.Context){ 
	var params v1dto.CreatePartDirectionInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := rh.service.CreatePartDirection(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (rh *ExamHandler) GetExamStructure(ctx *gin.Context){
	var params GetIdExamParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	examStructure, err := rh.service.GetExamStructure(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", examStructure)
}

func (rh *ExamHandler) GetExamPart(ctx *gin.Context){
	var params GetExamPartParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	examPart, err := rh.service.GetExamPart(params.ExamId, params.PartId)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", examPart)
}