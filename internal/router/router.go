package router

import (
	"shortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)

	api := r.Group("/api")

	api.POST("/shorten", handlers.CreateUrl)
}
