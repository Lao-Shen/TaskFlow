package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================
//  接口处理
// ============================================================

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
