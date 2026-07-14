package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "root:123456@tcp(192.168.188.129:3306)/TaskFlowData?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 确保数据库、连接、后续建表均使用 utf8mb4
	DB.Exec("SET NAMES utf8mb4")
	DB.Exec("ALTER DATABASE TaskFlowData CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")

	fmt.Println("数据库连接成功")
}

// SetTableCharset 强制指定表字符集（供 AutoMigrate 后调用）
func SetTableCharset(tableName string) {
	DB.Exec("ALTER TABLE " + tableName + " CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
}
