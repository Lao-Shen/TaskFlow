package user

import (
	"time"
)

// ============================================================
//  数据模型
// ============================================================

type User struct {
	ID        uint64    `gorm:"primaryKey"                   json:"id"`
	Username  string    `gorm:"size:32;uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"size:255;not null"            json:"-"`
	Email     string    `gorm:"size:128;uniqueIndex;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
