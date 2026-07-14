package task

import (
	"backend/pkg/database"

	"gorm.io/gorm"
)

// ============================================================
//  任务数据库操作
// ============================================================

func db() *gorm.DB {
	return database.DB
}

// AutoMigrate 自动迁移任务表
func AutoMigrate() error {
	return db().AutoMigrate(&Task{})
}

// ---------- 增 ----------

func InsertTask(task *Task) error {
	return db().Create(task).Error
}

// ---------- 查 ----------

func GetTaskByID(id uint64) (*Task, error) {
	var task Task
	err := db().First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetTasksByUserID(userID uint64, limit int) ([]Task, error) {
	var tasks []Task
	err := db().Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&tasks).Error
	return tasks, err
}

// ---------- 改 ----------

func UpdateTask(task *Task) error {
	return db().Save(task).Error
}

func UpdateTaskStatus(id uint64, status string) error {
	return db().Model(&Task{}).Where("id = ?", id).Update("status", status).Error
}

func UpdateTaskResult(id uint64, status, result, errMsg string) error {
	updates := map[string]any{
		"status":        status,
		"result":        result,
		"error_message": errMsg,
	}
	return db().Model(&Task{}).Where("id = ?", id).Updates(updates).Error
}

// GetTasksByUserIDAndStatus 按用户和可选状态查询任务列表
func GetTasksByUserIDAndStatus(userID uint64, status string, limit int) ([]Task, error) {
	q := db().Where("user_id = ?", userID)
	if status != "" {
		q = q.Where("status = ?", status)
	}
	var tasks []Task
	err := q.Order("created_at DESC").Limit(limit).Find(&tasks).Error
	return tasks, err
}

// GetTaskStatsByUserID 按用户统计各状态任务数量
func GetTaskStatsByUserID(userID uint64) (TaskStats, error) {
	var stats TaskStats
	stats.Total = db().Model(&Task{}).Where("user_id = ?", userID).Count(&stats.Total).RowsAffected

	// 各状态计数
	countByStatus := func(status string) int64 {
		var c int64
		db().Model(&Task{}).Where("user_id = ? AND status = ?", userID, status).Count(&c)
		return c
	}
	stats.Created = countByStatus(StatusCreated)
	stats.Queued = countByStatus(StatusQueued)
	stats.Running = countByStatus(StatusRunning)
	stats.Success = countByStatus(StatusSuccess)
	stats.Failed = countByStatus(StatusFailed)

	// 修正 total（首次 Count 返回值可能因 RowsAffected 不准确）
	stats.Total = stats.Created + stats.Queued + stats.Running + stats.Success + stats.Failed

	return stats, nil
}
