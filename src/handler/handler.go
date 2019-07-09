package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"net/http"
	"new_village/src/model"
	"new_village/src/service"
	"os"
	"strings"
)

type Result struct {
	STATUS    int
	Data      string
	ErrorDate string
}

func Create(c *gin.Context) {
	var server service.ServiceMan
	var req model.Order
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create BindJSON error", err)
	}

	_, err := server.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create Create error", err)
		return
	}

	b, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: fmt.Sprintf("解析结果失败 ,%v", err)})
		log.Print("Create Marshal error", err)
		return
	}

	c.JSON(http.StatusOK, Result{STATUS: 0, Data: string(b)})
}

func Update(c *gin.Context) {
	var server service.ServiceMan
	var req model.Order
	if err := c.BindJSON(&req); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create BindJSON error", err)
	}

	err := server.UpdateById(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create UpdateById error", err)
		return
	}
	c.JSON(http.StatusOK, Result{STATUS: 0, Data: "修改成功"})
}

func GetDetail(c *gin.Context) {
	var server service.ServiceMan
	orderId := c.Param("order_id")
	//去掉空格
	if len(strings.TrimSpace(orderId)) == 0 {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: "order_id is nil"})
		log.Print("order_id is nil")
		return
	}

	order, err := server.SelectOrderById(&model.Order{Order_id: orderId})
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create SelectOrderById error", err)
		return
	}

	//数据转换
	da, err := json.Marshal(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Result{STATUS: 0, Data: string(da)})
	//写入文件
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = string(da)
	err = file.Save("new_village.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

	//下载文件
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", "new_village.xlsx"))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./file/new_village.xlsx")
}

/**
上传文件
*/
func UploadingPath(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get file err:%s", err.Error()))
		return
	}

	//获取文件名称
	filename := header.Filename
	//写入文件
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	filepath := c.Param("fileUrl")
	//注册文件上传处理器
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func UploadingFile(c *gin.Context) {
	router := gin.Default()
	router.POST("/upload", UploadingPath)
	//注册静态文件路径
	router.StaticFS("/file", http.Dir("./public"))
}

func getCondition(c *gin.Context) {
	var server service.ServiceMan
	var req model.QueryCondition
	if err := c.BindJSON(&req); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create BindJSON error", err)
	}

	result, err := server.Select(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: err.Error()})
		log.Print("Create UpdateById error", err)
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, Result{STATUS: 1, ErrorDate: fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, Result{STATUS: 0, Data: string(res)})

}
