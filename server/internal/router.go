package internal

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.GET("/ping", Ping)

	api := r.Group("/api")

	api.POST("/", CreateUrl)
	api.GET("/:shortcode", RetrieveUrl)
	api.PUT("/:shortcode", UpdateUrl)
	api.DELETE("/:shortcode", UpdateUrl)
}
