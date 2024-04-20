package app

import (
	"context"

	"github.com/google/uuid"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	if repo == nil {
		panic("missing TaskRepository")
	}

	return &TaskService{repo: repo}
}

func (c TaskService) GetAllTasks(ctx context.Context) ([]Task, error) {
	return c.repo.GetAll(ctx)
}

func (c TaskService) CreateTask(ctx context.Context, title string, description string) error {
	task := Task{
		UUID:        uuid.New().String(),
		Title:       title,
		Description: description,
	}

	return c.repo.Create(ctx, &task)
}

func (c TaskService) DeleteTask(ctx context.Context, uuid string) error {
	return c.repo.Delete(ctx, uuid)
}

func (c TaskService) GetOneTask(ctx context.Context, uuid string) (*Task, error) {
	return c.repo.GetByID(ctx, uuid)
}

func (c TaskService) UpdateTask(ctx context.Context, uuid string, task *Task) error {
	return c.repo.Update(ctx, uuid, task)
}
