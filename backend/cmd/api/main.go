package main

import (
	"backend/internal/router"
	"backend/internal/task"
	"backend/internal/user"
	"backend/pkg/database"
	"fmt"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 自动迁移用户表、任务表
	if err := user.AutoMigrate(); err != nil {
		panic(fmt.Errorf("用户表迁移失败: %w", err))
	}
	database.SetTableCharset("users")

	if err := task.AutoMigrate(); err != nil {
		panic(fmt.Errorf("任务表迁移失败: %w", err))
	}
	database.SetTableCharset("tasks")

	// 启动 HTTP 服务
	r := router.ServerInit()

	fmt.Println("TaskFlow 服务启动: http://localhost:8080")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
