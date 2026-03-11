package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	v1service "github.com/dangLuan01/ets-api/internal/service/v1/skill"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	service v1service.SkillService
}

type EditSkillParams struct {
	Id int `uri:"id" binding:"required"`
}

func NewSkillHandler(service v1service.SkillService) *SkillHandler {
	return &SkillHandler {
		service: service,
	}
}

func (ch *SkillHandler) GetAllSkills(ctx *gin.Context) {
	skills, err := ch.service.GetAllSkills()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", skills)
}

func (ch *SkillHandler) CreateSkill(ctx *gin.Context) {
	var params v1dto.SkillParamsInput
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.CreateSkill(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *SkillHandler) EditSkill(ctx *gin.Context) {
	var params EditSkillParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	skill, err := ch.service.EditSkill(params.Id)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Successfully.", skill)
}

func (ch *SkillHandler) UpdateSkill(ctx *gin.Context) {
	var params v1dto.SkillParamsUpdate
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	if err := ch.service.UpdateSkill(params); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatus(ctx, http.StatusNoContent)
}

func (ch *SkillHandler) DeleteSkill(ctx *gin.Context) {}