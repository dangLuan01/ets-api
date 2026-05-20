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
	user := r.Group("/user")
	{
		user.GET("/info", ur.handler.GetUserByUUID)
		// users.POST("", ur.handler.CreateUser)
		user.PUT("/update", ur.handler.UpdateUser)
		// users.DELETE("/:uuid", ur.handler.DeleteUser)
		user.PATCH("/update-password", ur.handler.UpdatePassword)
		user.GET("/attempt", ur.handler.GetAttemptByUserUUID)
	}
}