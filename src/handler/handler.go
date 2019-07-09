package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"new_village/src/model"
	"new_village/src/service"
	"strings"
)

type Result struct {
	STATUS    int    `gorm:"status" json:"status"`       //状态：0表示成功， 1表示失败
	Data      string `gorm:"data" json:"data"`           //返回数据，状态为0时有值
	ErrorDate string `jorm:"error_str" json:"error_str"` //异常数据，状态为1时有值
	message   string `gorm:"message json:"message""`     //备注
}

func Create(c *gin.Context) {
	var server service.ServiceMan
	var req model.Order
	if err := c.BindJSON(&req); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: err.Error()})
	}

	_, err := server.Create(&req)
	if err != nil {
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: err.Error()})
		return
	}

	b, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: fmt.Sprintf("解析结果失败 ,%v", err)})
		return
	}

	c.JSON(http.StatusOK, Result{STATUS: 0, Data: string(b)})
}

func Update(c *gin.Context) {
	var server service.ServiceMan
	var req model.Order
	if err := c.BindJSON(&req); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: err.Error()})
	}

	err := server.UpdateById(&req)
	if err != nil {
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Result{STATUS: 0, Data: "修改成功"})
}

func GetDetail(c *gin.Context) {
	var server service.ServiceMan
	orderId := c.Param("order_id")
	//去掉空格
	if len(strings.TrimSpace(orderId)) == 0 {
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: "order_id is nil"})
		return
	}

	order, err := server.SelectOrderById(&model.Order{Order_id: orderId})
	if err != nil {
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: err.Error()})
		return
	}

	//数据转换
	da, err := json.Marshal(order)
	if err != nil {
		c.JSON(http.StatusOK, Result{STATUS: 1, ErrorDate: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Result{STATUS: 0, Data: string(da)})
}
