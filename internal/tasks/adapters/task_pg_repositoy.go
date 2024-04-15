package adapters

import (
	"context"

	"github.com/DuckyMomo20012/go-todo/internal/tasks/app"
	"xorm.io/xorm"
)

type PgTaskRepository struct {
	engine *xorm.Engine
}

func NewPgTaskRepository(engine *xorm.Engine) *PgTaskRepository {
	return &PgTaskRepository{engine: engine}
}

func (p PgTaskRepository) GetAll(_ context.Context) ([]app.Task, error) {
	var tasks []app.Task

	err := p.engine.Find(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p PgTaskRepository) Create(_ context.Context, task *app.Task) error {
	_, err := p.engine.Insert(task)

	return err
}

func (p PgTaskRepository) Delete(_ context.Context, id string) error {
	_, err := p.engine.ID(id).Delete(&app.Task{})

	return err
}

func (p PgTaskRepository) GetById(_ context.Context, id string) (*app.Task, error) {
	var task app.Task

	_, err := p.engine.ID(id).Get(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (p PgTaskRepository) Update(_ context.Context, id string, task *app.Task) error {
	_, err := p.engine.ID(id).Update(task)

	return err
}
