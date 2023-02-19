package cors

import (
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func CorsHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if origin := ctx.GetHeader("Origin"); origin != "" {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Header("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		// Stop here if its Preflighted OPTIONS request
		if ctx.Request.Method == "OPTIONS" {
			return
		}
		ctx.Next()
	}
}
