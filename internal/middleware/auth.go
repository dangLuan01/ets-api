package middleware

import (
	"net/http"
	"strings"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/pkg/auth"
	"github.com/dangLuan01/ets-api/pkg/cache"
	"github.com/gin-gonic/gin"
)

var (
	jwtService auth.TokenService
	cacheService cache.RedisCacheService
)

func InitAuthMiddlware(service auth.TokenService, cache cache.RedisCacheService) {
	jwtService = service
	cacheService = cache
}

func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		authHeder := ctx.GetHeader("Authorization")
		if authHeder == "" || !strings.HasPrefix(authHeder, "Bearer ") {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}

		tokenString := strings.TrimPrefix(authHeder, "Bearer ")

		_, claims, err := jwtService.ParseToken(tokenString)
		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}

		if jti, ok := claims["jti"].(string); ok {
			key := "blacklist:" + jti
			exists, err := cacheService.Exits(key)
			if err == nil && exists {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token revoked",
				})
				return 
			}
		}

		payload, err := jwtService.DecryptAccessTokenPayload(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}
		ctx.Set("data", payload)
		
		ctx.Next()
		
	}
}

func OptinalAuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		authHeder := ctx.GetHeader("Authorization")
		if authHeder == "" || !strings.HasPrefix(authHeder, "Bearer ") {

			ctx.Next()
			return 
		}

		tokenString := strings.TrimPrefix(authHeder, "Bearer ")

		_, claims, err := jwtService.ParseToken(tokenString)
		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}

		if jti, ok := claims["jti"].(string); ok {
			key := "blacklist:" + jti
			exists, err := cacheService.Exits(key)
			if err == nil && exists {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token revoked",
				})
				return 
			}
		}

		payload, err := jwtService.DecryptAccessTokenPayload(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}
		ctx.Set("data", payload)
		
		ctx.Next()	
	}
}

func RoleRequired(allowRoles int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, exists := ctx.Get("data")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}

		payload := val.(*v1dto.EncryptedPayload)
		isAllowed := false
		if payload.Role == int8(allowRoles) {
			isAllowed = true
		}

		if !isAllowed {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing or invalid.",
			})
			return 
		}

		ctx.Next()
	}
}