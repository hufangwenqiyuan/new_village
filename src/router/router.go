package router

import (
	"github.com/gin-gonic/gin"
	. "new_village/src/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//插入数据
	router.POST("/insert", Create)
	//更新数据
	router.PUT("/update", Update)
	//根据id获取数据
	router.GET("/select", GetDetail)
	//获取按条件排序的数据
	router.GET("/getCondition", getCondition)

	//上传文件接口
	router.GET("/uploading", UploadingFile)
	return router
}
