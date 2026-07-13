package internal

import (
	"backend/internal/database"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "root:123456@tcp(192.168.188.129:3306)/TaskFlowData?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	fmt.Println("数据库连接成功")
}