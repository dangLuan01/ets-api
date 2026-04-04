package v1routes

import (
	v1handler "github.com/dangLuan01/ets-api/internal/handler/v1/user"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *v1handler.UserHandler
}

func NewUserRoutes(handler *v1handler.UserHandler) *UserRoutes {
	return &UserRoutes {
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", ur.handler.GetAllUser)
		users.POST("/info", ur.handler.GetUserByUUID)
		users.POST("", ur.handler.CreateUser)
		users.PUT("/:uuid", ur.handler.UpdateUser)
		users.DELETE("/:uuid", ur.handler.DeleteUser)
		users.PUT("/change-password", ur.handler.ChangePassword)
	}
}