package task

import (
	"errors"
)

// ============================================================
//  业务错误定义
// ============================================================

var (
	ErrTaskCreateFailed = errors.New("任务创建失败")
)

// ============================================================
//  创建任务
// ============================================================

func createTask(userID uint64, req CreateTaskRequest) (*Task, error) {
	t := &Task{
		UserID:  userID,
		Name:    req.Name,
		Type:    req.Type,
		Status:  StatusCreated,
		Payload: req.Payload,
	}

	if err := InsertTask(t); err != nil {
		return nil, ErrTaskCreateFailed
	}

	return t, nil
}

// ============================================================
//  查询任务列表
// ============================================================

func listTasks(userID uint64, status string, limit int) (*TaskListResponse, error) {
	if limit <= 0 || limit > 200 {
		limit = 200
	}

	tasks, err := GetTasksByUserIDAndStatus(userID, status, limit)
	if err != nil {
		return nil, err
	}

	stats, err := GetTaskStatsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 转换响应
	responses := make([]TaskResponse, len(tasks))
	for i, t := range tasks {
		responses[i] = TaskResponse{
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
		}
	}

	return &TaskListResponse{Tasks: responses, Stats: stats}, nil
}
