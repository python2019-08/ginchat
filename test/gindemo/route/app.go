package route

import (
	"ginchat/test/gindemo/app/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.GET("/info", controller.InfoHandle())
	r.POST("/add", controller.AddHandle())
	r.POST("/edit", controller.EditHandle())
	r.POST("/del", controller.DelHandle())
}
