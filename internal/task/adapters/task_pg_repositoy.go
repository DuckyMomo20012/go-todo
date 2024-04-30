package adapters

import (
	"context"

	"github.com/DuckyMomo20012/go-todo/internal/task/app"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

type PgTaskRepository struct {
	db *pgxpool.Pool
}

func NewPgTaskRepository(db *pgxpool.Pool) *PgTaskRepository {
	if db == nil {
		panic("missing db connection")
	}

	return &PgTaskRepository{
		db: db,
	}
}

func (p PgTaskRepository) CreateTask(ctx context.Context, body *app.CreateTaskDto) (*app.Task, error) {
	q, args := psql.Insert(
		im.Into("task", "title", "description"),
		im.Values(psql.Raw("nullif(?, '')", body.Title), psql.Arg(body.Description)),
		im.Returning("*"),
	).MustBuild()

	rows, err := p.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	createdTask, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[app.Task])
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (p PgTaskRepository) GetAllTask(ctx context.Context) ([]*app.Task, error) {
	q, args := psql.Select(
		sm.Columns("*"),
		sm.From("task"),
	).MustBuild()

	rows, err := p.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[app.Task])
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p PgTaskRepository) GetTaskById(ctx context.Context, taskId string) (*app.Task, error) {
	q, args := psql.Select(
		sm.Columns("*"),
		sm.From("task"),
		sm.Where(psql.Quote("task_id").EQ(psql.Arg(taskId))),
	).MustBuild()

	rows, err := p.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	task, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[app.Task])
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (p PgTaskRepository) UpdateTask(ctx context.Context, taskId string, body *app.UpdateTaskDto) (*app.Task, error) {
	q, args := psql.Update(
		um.Table("task"),
		um.SetCol("title").To(psql.Raw("coalesce(nullif(?, ''), title)", body.Title)),
		um.SetCol("description").To(psql.Raw("coalesce(?, description)", body.Description)),
		um.Where(psql.Quote("task_id").EQ(psql.Arg(taskId))),
		um.Returning("*"),
	).MustBuild()

	rows, err := p.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	updatedTask, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[app.Task])
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (p PgTaskRepository) DeleteTask(ctx context.Context, taskId string) (*app.Task, error) {
	q, args := psql.Delete(
		dm.From("task"),
		dm.Where(psql.Quote("task_id").EQ(psql.Arg(taskId))),
		dm.Returning("*"),
	).MustBuild()

	rows, err := p.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	deletedTask, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[app.Task])
	if err != nil {
		return nil, err
	}

	return deletedTask, nil
}
