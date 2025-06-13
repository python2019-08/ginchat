package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InfoHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里从模型获取数据
		// fmt.Println("InfoHandle() gin.HandlerFunc")
		c.String(http.StatusOK, "hello, info")
	}
}

func AddHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里向模型添加数据
		fmt.Println("func AddHandle() gin.HandlerFunc {")
		c.String(http.StatusOK, "hello, AddHandle")
	}
}

func EditHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里向模型编辑数据
		fmt.Println("func EditHandle() gin.HandlerFunc {")
		c.String(http.StatusOK, "hello, EditHandle")
	}
}

func DelHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里从模型删除数据
		fmt.Println("func DelHandle() gin.HandlerFunc {")
		c.String(http.StatusOK, "hello, DelHandle")
	}
}
