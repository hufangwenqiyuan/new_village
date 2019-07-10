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
	Amount string `JSON:"amount" form:"amount"`
	//状态
	Status string `JSON:"status" form:"status"`
	//文件路径
	File_url string `JSON:"file_url" form:"file_url"`
}

type QueryCondition struct {
	Key     string
	LikeStr string
	Desc    bool
}
