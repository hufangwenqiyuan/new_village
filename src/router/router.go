package router

import (
	"github.com/gin-gonic/gin"
	. "new_village/src/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/insert", Create)
	router.PUT("/update", Update)
	router.GET("/select", GetDetail)
	return router
}
