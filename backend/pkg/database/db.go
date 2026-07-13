package database

import "gorm.io/gorm"

// DB 全局数据库实例，由 internal.InitDB() 初始化
var DB *gorm.DB