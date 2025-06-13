package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// http://localhost:8081/index
	r.GET("/index", service.GetIndex)
	return r
}
