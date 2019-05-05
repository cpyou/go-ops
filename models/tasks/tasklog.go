package tasks

import (
	"database/sql"
	"go-ops/models"
)

type TaskLog struct {
	models.BaseModel
	Description sql.NullString `json:"description"`
	Action      string         `json:"action"`      // run stop
	UserId      uint           `json:"user_id"`
	TaskId      uint           `json:"task_id"`
}
