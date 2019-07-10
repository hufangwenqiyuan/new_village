package service

import (
	"fmt"
	"log"
	"new_village/src/model"
	"testing"
)

func TestCreate(t *testing.T) {
	//d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	//if err != nil {
	//	log.Fatalf("open the DB fail, err:%s", err)
	//}
	id, err := ServiceManage().Create(&model.Order{
		ID:        12,
		Order_id:  "9876",
		Amount:    "234",
		File_url:  "src",
		User_name: "hufangwen",
		Status:    "正常",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("text mode Create date", id)
}

func TestUpdateById(t *testing.T) {
	//d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	//if err != nil {
	//	log.Fatalf("open the DB fail, err:%s", err)
	//}
	err := ServiceManage().UpdateById(&model.Order{
		Order_id:  "0987",
		Amount:    "73",
		File_url:  "src",
		User_name: "hufangwen",
		Status:    "正常",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestSelectOrderById(t *testing.T) {
	//d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	//if err != nil {
	//	log.Fatalf("open the DB fail, err:%s", err)
	//}

	ret, err := ServiceManage().SelectOrderById(&model.Order{
		Order_id: "1231"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Println("%v", ret))
}

func TestSelect(t *testing.T) {
	//d, err := gorm.Open("mysql", "root:root/test?charset=utf8")
	//if err != nil {
	//	log.Fatalf("open the DB fail, err:%s", err)
	//}
	ret, erre := ServiceManage().Select(&model.QueryCondition{
		LikeStr: "1231",
		Desc:    true,
	})

	if erre != nil {
		log.Fatal(erre)
	}
	fmt.Println(fmt.Println("%v", ret))
}
