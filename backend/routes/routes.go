package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/toybox97/go-sse/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheck)
}