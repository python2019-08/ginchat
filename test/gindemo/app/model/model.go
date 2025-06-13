package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Phone string `json:"phone"`
}

// 设置表名
func (User) TableName() string {
	return "user"
}
