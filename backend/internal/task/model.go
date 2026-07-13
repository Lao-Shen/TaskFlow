package task

import "time"

type Task struct {
	ID uint64 `gorm:"primaryKey"`

	UserID uint64 `gorm:"not null;index"`

	Name string `gorm:"size:100;not null"`

	Type string `gorm:"size:50;not null;index"`

	Status string `gorm:"size:20;not null;index"`

	Payload string `gorm:"type:text"`

	Result string `gorm:"type:text"`

	RetryCount int `gorm:"default:0"`

	ErrorMessage string `gorm:"type:text"`

	CreatedAt time.Time

	UpdatedAt time.Time
}

const (
	StatusCreated = "CREATED"
	StatusQueued  = "QUEUED"
	StatusRunning = "RUNNING"
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)
