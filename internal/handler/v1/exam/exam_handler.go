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

type GetSlugExamParams struct {
	Slug string `uri:"slug" binding:"required"`
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

func (eh *ExamHandler) FindExamBySlug(ctx *gin.Context) {
	var params GetSlugExamParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	exam, err := eh.service.FindExamBySlug(params.Slug)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", exam)
}

func (eh *ExamHandler) CalculateScoreExam(ctx *gin.Context) {
	var params v1dto.QuestionAnswerInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	data, err := eh.service.CalculateScoreExam(ctx, params)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", v1dto.MapDetailExamScoreDTO(data))
}

func (eh *ExamHandler) GetFeatured(ctx *gin.Context) {
	var params v1dto.ExamFeaturedParams
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

	exams, totalRecords, err := eh.service.GetFeaturedExams(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, exams))
}

// --- HANDLER CHO ADMIN (CRUD EXAM) ---

func (eh *ExamHandler) GetAllExams(ctx *gin.Context){
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

	exams, totalRecords, err := eh.service.GetAllExams(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	paginationResponse := utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, exams)
	
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", paginationResponse)
}

func (eh *ExamHandler) CreateExam(ctx *gin.Context){
	var params v1dto.CreateExamInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := eh.service.CreateExam(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) EditExam(ctx *gin.Context){
	var params GetIdExamParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	exam, err := eh.service.EditExamById(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", exam)
}

func (eh *ExamHandler) UpdateExam(ctx *gin.Context){
	var params v1dto.UpdateExamInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}
	
	if err := eh.service.UpdateExam(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) CreatePartDirection(ctx *gin.Context){ 
	var params v1dto.CreatePartDirectionInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := eh.service.CreatePartDirection(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) GetExamStructure(ctx *gin.Context){
	var params GetIdExamParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	examStructure, err := eh.service.GetExamStructure(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", examStructure)
}

func (eh *ExamHandler) GetExamPart(ctx *gin.Context){
	var params GetExamPartParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	examPart, err := eh.service.GetExamPart(params.ExamId, params.PartId)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", examPart)
}

func (eh *ExamHandler) UpdateQuestionSingle(ctx *gin.Context){
	var params v1dto.UpdateQuestionSingleInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := eh.service.UpdateQuestionSingle(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) UpdateQuestionGroup(ctx *gin.Context){
	var params v1dto.UpdateQuestionGroupInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := eh.service.UpdateQuestionGroup(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) UpdatePartDirection(ctx *gin.Context){
	var params v1dto.UpdatePartDirectionInputParams
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := eh.service.UpdatePartDirection(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) ImportExamQuestionFromExcel(ctx *gin.Context){
	var params v1dto.ExamImportInputParams
	if err := ctx.ShouldBind(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := eh.service.ImportExamQuestionFromExcel(ctx, params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (eh *ExamHandler) GetFilterStructure(ctx *gin.Context) {
	
	filterStructre, err := eh.service.GetFilterStructure()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	
	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", filterStructre)
}

func (eh *ExamHandler) FilterExam(ctx *gin.Context) {
	var params v1dto.FilterExamParams
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

	exams, totalRecords, err := eh.service.FilterExam(params)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", utils.NewPaginationResponse(params.Page, params.Limit, totalRecords, exams))
}