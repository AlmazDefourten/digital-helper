package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/chat/send", v1.GetAnswer)

	return r
}
