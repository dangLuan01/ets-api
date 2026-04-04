package routes

import (
	"net/http"

	"github.com/dangLuan01/ets-api/internal/middleware"
	v1routes "github.com/dangLuan01/ets-api/internal/routes/v1"
	v1routesClient "github.com/dangLuan01/ets-api/internal/routes/v1/client"
	"github.com/dangLuan01/ets-api/pkg/auth"
	"github.com/dangLuan01/ets-api/pkg/cache"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoute(r *gin.Engine, authService auth.TokenService, cacheService cache.RedisCacheService , routes ...Route) {
	r.Use(
		middleware.CORSMiddleware(),
	)
	
	v1api 		:= r.Group("/api/v1")
	protected 	:= v1api.Group("")

	v1api.Use(
		//middleware.ApiKeyMiddleware(),
		middleware.OptinalAuthMiddleware(),
		middleware.RateLimiterMiddleware(),
	)
	
	middleware.InitAuthMiddlware(authService, cacheService)
	
	protected.Use(
		
		middleware.AuthMiddleware(),
		middleware.RoleRequired(1),
		//middleware.ApiKeyMiddleware(),
		middleware.RateLimiterMiddleware(),
	)

	for _, route := range routes {

		switch route.(type) {
		case *v1routes.AuthRoutes:
			route.Register(v1api)
		case *v1routesClient.MenuRoutes:
			route.Register(v1api)
		case *v1routesClient.ExamRoutes:
			route.Register(v1api)
		default:
			route.Register(protected)
		}
	}

	r.NoRoute(func(ctx *gin.Context) {

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.JSON(404, gin.H{
			"error":"NOT FOUND",
			"path": ctx.Request.URL.Path,
		})
	})
}