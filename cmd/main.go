package main

import (
	"shortener/internal/config"
	"shortener/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	router.SetRouter(r)
	r.Run()
}
