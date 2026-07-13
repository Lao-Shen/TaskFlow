package main

import (
	"backend/internal/router"
	"backend/internal/user"
	"backend/pkg/database"
	"fmt"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 自动迁移用户表
	if err := user.AutoMigrate(); err != nil {
		panic(fmt.Errorf("数据库迁移失败: %w", err))
	}

	// 启动服务
	r := router.ServerInit()

	fmt.Println("TaskFlow 服务启动: http://localhost:8080")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
