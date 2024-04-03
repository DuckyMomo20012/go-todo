package app

import "context"

type TaskRepository interface {
	GetAll(ctx context.Context) ([]Task, error)
	Create(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, id string, task *Task) error
}
