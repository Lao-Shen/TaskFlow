package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================
//  接口处理
// ============================================================

// POST /api/register
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数校验失败: " + err.Error()})
		return
	}

	resp, err := registerUser(req)
	if err != nil {
		switch {
		case errors.Is(err, ErrUsernameTaken):
			c.JSON(http.StatusConflict, gin.H{"code": 409, "msg": err.Error()})
		case errors.Is(err, ErrEmailTaken):
			c.JSON(http.StatusConflict, gin.H{"code": 409, "msg": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "注册成功",
		"data": resp,
	})
}

// POST /api/login
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数校验失败: " + err.Error()})
		return
	}

	resp, err := loginUser(req)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": err.Error()})
		case errors.Is(err, ErrWrongPassword):
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": resp,
	})
}

// POST /api/logout
func Logout(c *gin.Context) {
	// JWT 无状态，登出由客户端删除 token 完成；服务端返回成功即可
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "已登出",
	})
}

// GET /api/user/me
func GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")

	id, ok := userID.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "用户信息异常"})
		return
	}

	user, err := GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": UserInfoResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}