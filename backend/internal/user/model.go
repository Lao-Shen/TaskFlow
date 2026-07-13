package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret 签名密钥 —— 生产环境应放入配置文件或环境变量
const JWTSecret = "TaskFlow-Secret-Key-1783914343"

// Claims 自定义 JWT 载荷
type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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

// ============================================================
//  请求 / 响应 DTO
// ============================================================

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Email    string `json:"email"    binding:"required,email,max=128"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UserInfoResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
