package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// http://localhost:8081/index
	r.GET("/index", service.GetIndex)
	// http://localhost:8081/user/getUserList
	r.GET("/user/getUserList", service.GetUserList)
	return r
}
