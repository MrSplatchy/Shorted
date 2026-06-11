package main

import (
	"shortener/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	internal.ConnectDB()
	internal.SetRouter(r)

	r.Run("0.0.0.0:8080")
}
