package main

import (
	"shortener/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	internal.ConnectDB()
	r := gin.Default()
	internal.SetRouter(r)
	r.Run()
}
