package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/skill"
	"github.com/gin-gonic/gin"
)

type SkillRoutes struct {
	handler *v1handler.SkillHandler
}

func NewSkillRoutes(handler *v1handler.SkillHandler) *SkillRoutes {
	return &SkillRoutes {
		handler: handler,
	}
}

func (cr *SkillRoutes) Register(r *gin.RouterGroup) {
	skill := r.Group("/skills")
	{
		skill.GET("/get-all", cr.handler.GetAllSkills)
		skill.POST("/create", cr.handler.CreateSkill)
		skill.GET("/edit/:id", cr.handler.EditSkill)
		skill.PUT("/update", cr.handler.UpdateSkill)
		skill.DELETE("/delete/:id", cr.handler.DeleteSkill)
	}
}