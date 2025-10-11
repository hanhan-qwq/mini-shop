package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mini_shop/config"
	"mini_shop/model"
)

var DBClient *gorm.DB

func InitMysql() {
	mysqlConfig := config.AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	DBClient = client
	log.Println("连接数据库成功")

	// 自动创建表（根据model下定义模型）
	if err := client.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("创建 users 表失败: ", err)
	}
	if err := client.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalln("创建 products 表失败：", err)
	}

}

func GetDB() *gorm.DB {
	return DBClient
}
func CloseDB() {
	if DBClient != nil {
		sqlDB, err := DBClient.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
