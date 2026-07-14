package task

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================
//  接口处理
// ============================================================

// POST /api/tasks  ——  创建任务
func CreateTask(c *gin.Context) {
	// 从 JWT 中间件注入的上下文中获取用户 ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未登录"})
		return
	}

	uid, ok := userID.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "用户信息异常"})
		return
	}

	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数校验失败: " + err.Error()})
		return
	}

	t, err := createTask(uid, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrTaskCreateFailed):
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器错误"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "任务创建成功",
		"data": TaskResponse{
			ID:           t.ID,
			UserID:       t.UserID,
			Name:         t.Name,
			Type:         t.Type,
			Status:       t.Status,
			Payload:      t.Payload,
			Result:       t.Result,
			RetryCount:   t.RetryCount,
			ErrorMessage: t.ErrorMessage,
			CreatedAt:    t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    t.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

// GET /api/tasks  ——  获取任务列表（含统计）
func ListTasks(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未登录"})
		return
	}

	uid, ok := userID.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "用户信息异常"})
		return
	}

	status := c.Query("status")
	resp, err := listTasks(uid, status, 200)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": resp,
	})
}
