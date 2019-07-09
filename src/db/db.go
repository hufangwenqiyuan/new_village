package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"new_village/src/model"
	"strings"
)

//数据库配置， GORM配置
//type db gorm.DB 行不通
var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:root:new_village?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	/**
	判断是否链接正常
	*/
	if err != nil {
		fmt.Printf("mysql connect err %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}

	database := Eloquent.AutoMigrate(&model.Order{})
	if database.Error != nil {
		fmt.Printf("create database error %v", database.Error)
	}

}

func NewManager(d *gorm.DB) *Sql {
	return &Sql{
		db: d,
	}
}

type Sql struct {
	db *gorm.DB
}

//增
func (db *Sql) Create(order *model.Order) (resultOrder *model.Order, err error) {
	//判断参数
	if order == nil {
		err = fmt.Errorf("order is nil")
		return
	}

	if strings.TrimSpace(order.Order_id) == "" || len(order.Order_id) > 40 {
		err = fmt.Errorf("Order_id is error,Order_id %v", order.Order_id)
		return
	}

	if strings.TrimSpace(order.User_name) == "" || len(order.User_name) > 40 {
		err = fmt.Errorf("User_name is error,User_name %v", order.User_name)
		return
	}
	if strings.TrimSpace(order.Amount) == "" || len(order.Amount) > 40 {
		err = fmt.Errorf("Amount is error,Amount %v", order.Amount)
		return
	}
	if strings.TrimSpace(order.Status) == "" || len(order.Status) > 40 {
		err = fmt.Errorf("Status is error,Status %v", order.Status)
		return
	}
	if strings.TrimSpace(order.File_url) == "" || len(order.File_url) > 300 {
		err = fmt.Errorf("File_url is error,Status %v", order.File_url)
		return
	}

	//开启事物
	thran := db.db.Begin()
	//判断是否需要回滚
	defer func() {
		if err != nil {
			thran.Rollback()
		} else {
			thran.Commit()
		}
	}()

	//条件注入，判断是否报错
	if !thran.Where("order_id=?", order.Order_id).RecordNotFound() {
		//指定需要运行的表，获取db，
		if er := thran.Table("new_village").Where("order_id=?", order.Order_id).Delete(&model.Order{}).Error; er != nil {
			err = er
			return
		}
	}

	//检查主键是否为空
	b := thran.NewRecord(&order)
	if !b {
		err = fmt.Errorf("check if value's primary key is blank return false")
		return
	}

	//插入数据库
	d := thran.Create(&order)

	if d.Error != nil {
		err = d.Error
		return
	}

	//检查主键是否为空
	b = thran.NewRecord(&order)
	if !b {
		err = fmt.Errorf("insert is fail")
		return
	}

	if order.ID < 1 {
		err = fmt.Errorf("return id is:%v", order.ID)
		return
	}
	return order, nil

}

//改

func (db *Sql) updateById(order *model.Order) (err error) {

	//判断参数
	if order == nil {
		err = fmt.Errorf("order is nil")
		return
	}

	if strings.TrimSpace(order.Order_id) == "" || len(order.Order_id) > 40 {
		err = fmt.Errorf("Order_id is error,Order_id %v", order.Order_id)
		return
	}

	if strings.TrimSpace(order.User_name) == "" || len(order.User_name) > 40 {
		err = fmt.Errorf("User_name is error,User_name %v", order.User_name)
		return
	}
	if strings.TrimSpace(order.Amount) == "" || len(order.Amount) > 40 {
		err = fmt.Errorf("Amount is error,Amount %v", order.Amount)
		return
	}
	if strings.TrimSpace(order.Status) == "" || len(order.Status) > 40 {
		err = fmt.Errorf("Status is error,Status %v", order.Status)
		return
	}
	if strings.TrimSpace(order.File_url) == "" || len(order.File_url) > 300 {
		err = fmt.Errorf("File_url is error,Status %v", order.File_url)
		return
	}

	//开启事物
	th := db.db.Begin()

	defer func() {
		if er := recover(); er != nil {
			th.Rollback()
		}
	}()

	if err := th.Model(&order).Select("amount", "status", "file_url").Updates(order).Error; err != nil {
		th.Rollback()
		return err
	}

	return nil

}

//根据要求查询数据
func (db *Sql) Select(condition *model.QueryCondition) ([]*model.Order, error) {
	if condition == nil {
		return nil, fmt.Errorf("condition is nil")
	}

	whereKey := ""
	whereValue := ""
	whereFlag := false
	if len(strings.TrimSpace(condition.LikeStr)) > 0 {
		whereKey = fmt.Sprintf("%s like ?", condition.Key)
		whereValue = "%" + condition.LikeStr + "%"
		whereFlag = true
	}
	if len(strings.TrimSpace(condition.LikeStr)) == 0 {
		condition.Key = "user_name"
	}

	desc := ""
	if condition.Desc {
		desc = "DESC"
	}

	checkOrder := make([]*model.Order, 0)
	th := &gorm.DB{}
	if whereFlag {
		th = db.db.Where(whereKey, whereValue).Order("amount " + desc).Order("create_time " + desc).Find(&checkOrder)
	} else {
		th = db.db.Order("amount " + desc).Order("create_time " + desc).Find(&checkOrder)
	}

	if th.Error != nil {
		return nil, th.Error
	}

	return checkOrder, nil
}

func (db *Sql) SelectOrderById(order *model.Order) (*model.Order, error) {
	if order == nil {
		return nil, fmt.Errorf("order is nil")
	}

	data := &model.Order{}
	th := db.db.Where("order_id=?", order.Order_id).Find(&data)
	if th.RecordNotFound() {
		return nil, fmt.Errorf("没有查询到相关信息 %v", order.Order_id)
	}
	if th.Error != nil {
		return nil, th.Error
	}
	return data, nil

}
