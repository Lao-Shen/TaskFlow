package task

// ============================================================
//  请求 / 响应 DTO
// ============================================================

type CreateTaskRequest struct {
	Name    string `json:"name"    binding:"required,min=1,max=100"`
	Type    string `json:"type"    binding:"required,min=1,max=50"`
	Payload string `json:"payload"`
}

type TaskResponse struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"user_id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Payload      string `json:"payload"`
	Result       string `json:"result"`
	RetryCount   int    `json:"retry_count"`
	ErrorMessage string `json:"error_message"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
	Stats TaskStats      `json:"stats"`
}

type TaskStats struct {
	Total   int64 `json:"total"`
	Created int64 `json:"created"`
	Queued  int64 `json:"queued"`
	Running int64 `json:"running"`
	Success int64 `json:"success"`
	Failed  int64 `json:"failed"`
}
