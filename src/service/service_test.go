package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"new_village/src/db"
	"new_village/src/model"
	"testing"
)

func TestCreat(t testing.T) {

	d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	if err != nil {
		log.Fatalf("open the DB fail, err:%s", err)
	}
	id, erre := ServiceManage(db.NewManager(d)).Create(&model.Order{
		Order_id:  "9876",
		Amount:    "23",
		File_url:  "src",
		User_name: "hufangwen",
		Status:    "正常",
	})
	if erre != nil {
		log.Fatal(err)
	}

	fmt.Println("text mode Create date", id)
}

func TestUpdateById(t *testing.T) {
	d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	if err != nil {
		log.Fatalf("open the DB fail, err:%s", err)
	}
	erre := ServiceManage(db.NewManager(d)).UpdateById(&model.Order{
		Order_id:  "0987",
		Amount:    "73",
		File_url:  "src",
		User_name: "hufangwen",
		Status:    "正常",
	})
	if erre != nil {
		log.Fatal(err)
	}
}

func TestSelect(t *testing.T) {
	d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	if err != nil {
		log.Fatalf("open the DB fail, err:%s", err)
	}

	ret, erre := ServiceManage(db.NewManager(d)).SelectOrderById(&model.Order{
		Order_id: "1231"})
	if erre != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Println("%v", ret))
}

func TestSelect(t *testing.T) {
	d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	if err != nil {
		log.Fatalf("open the DB fail, err:%s", err)
	}
	ret, erre := ServiceManage(db.NewManager(d)).Select(&model.QueryCondition{
		LikeStr: "1231",
		Desc:    true,
	})

	if erre != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Println("%v", ret))
}
