package app

import "github.com/jackc/pgx/v5/pgtype"

type Task struct {
	TaskId      string           `json:"taskId"`
	Title       string           `json:"title"`
	Description *string          `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
}

type CreateTaskDto struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}

type UpdateTaskDto struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
