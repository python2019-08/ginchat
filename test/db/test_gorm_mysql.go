package dbtest

// https://gorm.io/zh_CN/docs/index.html
//
// https://github.com/go-gorm/gorm
// https://gorm.io/docs/
import (
	"fmt"
	"ginchat/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_gorm_mysql() {
	var dsn = "root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		log.Fatalln("db.AutoMigrate() error:", err)
	}

	user := &models.UserBasic{}
	user.Name = "申专"
	user.LoginTime = time.Now()
	user.HeartbeatTime = time.Now()
	user.LoginOutTime = time.Now()

	result := db.Create(user)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	fmt.Println("CreateRecord():", result.RowsAffected, result.Error, user)

	//Read
	user1 := &models.UserBasic{}
	result = db.Unscoped().First(user1, "1") //根据整型主键查找
	if result.Error != nil {
		log.Println("db.Unscoped().First() err=", result.Error)
		return
	}
	fmt.Println("db.Unscoped().First():RowsAffected", result.RowsAffected,
		".......result.Error=", result.Error,
		".......user1=", user1)

	//db.First（user，"code=?"，"D42"）//查找code字段值为D42的记录

	// Update - 将product的price更新为2oo
	result = db.Model(user).Update("PassWord", "1234")
	if result.Error != nil {
		log.Println("db.Model(user).Update() err=", result.Error)
		return
	}
	fmt.Println("db.Model(user).Update(PassWord=1234):RowsAffected", result.RowsAffected,
		".......result.Error=", result.Error)

	//Update-更新多个子段
	//db.Model（&product）.Updates（Product{Price：2o0，Code："F42"}）//仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}f"Price":200,"Code":"F42"})
	//Delete-删除product
	//db.Delete(&product,1)
}
