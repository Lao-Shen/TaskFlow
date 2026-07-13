package authjwt

import "github.com/golang-jwt/jwt/v5"

// JWTSecret 签名密钥 —— 生产环境应放入配置文件或环境变量
const JWTSecret = "TaskFlow-Secret-Key-1783914343"

// Claims 自定义 JWT 载荷
type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
