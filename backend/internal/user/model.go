package user

import (
	"time"
)

// 数据结构

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"size:32;unique"`
	Password  string `gorm:"size:255"`
	Email     string `gorm:"size:128"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
