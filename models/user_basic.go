package models

import (
	// "github.com/jinzhu/gorm" // old-version gorm
	"time"

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
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

// 在 GORM 中，`TableName()` 方法是一个**表名钩子（Hook）**，它的作用是**自定义数据库表名**。
// 当你调用 `AutoMigrate` 时，GORM 会根据这个方法的返回值来创建或更新对应的数据库表。
func (table *UserBasic) TableName() string {
	return "user_basic"
}
