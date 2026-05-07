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
	
	//v1api 			:= r.Group("/api/v1")
	// authRoute   	:= v1api.Group("")
	// protected 		:= v1api.Group("")
	// protectedClient := v1api.Group("")

	// v1api.Use(
	// 	//middleware.ApiKeyMiddleware(),
	// 	middleware.OptinalAuthMiddleware(),
	// 	middleware.RateLimiterMiddleware(),
	// )

	// authRoute.Use(
	// 	middleware.RateLimiterMiddleware(),
	// )
	
	// middleware.InitAuthMiddlware(authService, cacheService)
	// protected.Use(
		
	// 	middleware.AuthMiddleware(),
	// 	middleware.RoleRequired(2),
	// 	//middleware.ApiKeyMiddleware(),
	// 	middleware.RateLimiterMiddleware(),
	// )

	// for _, route := range routes {

	// 	switch route.(type) {
	// 	case *v1routes.AuthRoutes:
	// 		route.Register(authRoute)
	// 	case *v1routesClient.MenuRoutes:
	// 		route.Register(v1api)
	// 	case *v1routesClient.ExamRoutes:
	// 		route.Register(v1api)
	// 	case *v1routesClient.PostRoutes:
	// 		route.Register(v1api)
	// 	case *v1routesClient.TagRoutes:
	// 		route.Register(v1api)
	// 	case *v1routesClient.UserRoutes:
	// 		route.Register(v1api)
	// 	default:
	// 		route.Register(protected)
	// 	}
	// }
	v1api := r.Group("/api/v1")

	// Global middleware cho toàn bộ API
	v1api.Use(
		middleware.RateLimiterMiddleware(),
	)

	// Init auth dependency
	middleware.InitAuthMiddlware(authService, cacheService)
	
	// =====================
	// Public routes
	// =====================
	public := v1api.Group("")
	public.Use(
		middleware.OptionalAuthMiddleware(),
	)

	// =====================
	// Auth routes
	// =====================
	auth := v1api.Group("")
	auth.Use(
		middleware.RateLimiterMiddleware(),
	)

	// =====================
	// Protected user routes
	// =====================
	client := v1api.Group("")
	client.Use(
		middleware.AuthMiddleware(),
		middleware.RoleRequired(2),
	)

	// =====================
	// Protected admin routes
	// =====================
	admin := v1api.Group("")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.RoleRequired(1),
	)

	// =====================
	// Register routes
	// =====================
	for _, route := range routes {
		switch route.(type) {
		case *v1routes.AuthRoutes:
			route.Register(auth)
		case *v1routesClient.MenuRoutes:
			route.Register(public)
		case *v1routesClient.ExamRoutes:
			route.Register(public)
		case *v1routesClient.PostRoutes:
			route.Register(public)
		case *v1routesClient.TagRoutes:
			route.Register(public)
		case *v1routesClient.UserRoutes:
			route.Register(client)
		// Default = admin protected
		default:
			route.Register(admin)
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