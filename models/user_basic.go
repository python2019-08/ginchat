package models

import (
	// "github.com/jinzhu/gorm" // old-version gorm
	"gorm.io/gorm" // new-version gorm
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogOutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

// 在 GORM 中，`TableName()` 方法是一个**表名钩子（Hook）**，它的作用是**自定义数据库表名**。
// 当你调用 `AutoMigrate` 时，GORM 会根据这个方法的返回值来创建或更新对应的数据库表。
func (table *UserBasic) TableName() string {
	return "user_basic"
}
