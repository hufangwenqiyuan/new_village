package model

/**
pojo实体类
*/
type Order struct {
	//唯一标识
	ID int64 `gorm:"primary_key" JSON：“id”`
	//顺序标识
	Order_id string `JSON:"order_id"`
	//用户名称
	User_name string `JSON:"user_name"`
	//总量
	Amount string `JSON:"amount"`
	//状态
	Status string `JSON:"status"`
	//文件路径
	File_url string `JSON:"file_url`
}

type QueryCondition struct {
	//查询字段
	Key string
	//关键字段
	LikeStr string
	//是否倒序
	Desc bool
	//页码
	Page int
	//每页的页码
	pageSize int
}
