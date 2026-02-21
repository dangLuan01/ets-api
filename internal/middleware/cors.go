package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

        ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Requested-With, x-api-key, Range")
        ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
        ctx.Header("Access-Control-Max-Age", "86400")
        ctx.Header("Access-Control-Expose-Headers", "Content-Length, Content-Range")

        if ctx.Request.Method == http.MethodOptions {
            ctx.AbortWithStatus(http.StatusNoContent)
            return
        }

		ctx.Next()
	}
}