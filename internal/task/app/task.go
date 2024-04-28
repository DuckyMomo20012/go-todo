package app

import "github.com/jackc/pgx/v5/pgtype"

type Task struct {
	TaskId      string  `json:"task_id" xorm:"varchar(255) not null pk 'task_id'"`
	Title       string  `json:"title" xorm:"varchar(255) not null 'title'"`
	Description *string `json:"description" xorm:"varchar(255) not null 'description'"`
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type CreateTaskDto struct {
	Title       string
	Description *string
}

type UpdateTaskDto struct {
	Title       *string
	Description *string
}
