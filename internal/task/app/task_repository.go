package app

import "context"

type TaskRepository interface {
	CreateTask(ctx context.Context, body *CreateTaskDto) (*Task, error)
	GetAllTask(ctx context.Context) ([]*Task, error)
	GetTaskById(ctx context.Context, taskId string) (*Task, error)
	UpdateTask(ctx context.Context, taskId string, body *UpdateTaskDto) (*Task, error)
	DeleteTask(ctx context.Context, taskId string) (*Task, error)
}
