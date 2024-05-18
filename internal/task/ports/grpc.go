package ports

import (
	"context"
	"errors"
	"fmt"

	taskv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/task/v1"
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/validate"
	"github.com/DuckyMomo20012/go-todo/internal/task/app"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	taskRepo app.TaskRepository
	taskv1.UnimplementedTaskServiceServer
}

func NewGrpcServer(taskRepo app.TaskRepository) GrpcServer {
	log := logger.Get()

	if taskRepo == nil {
		log.Panic().Msg("missing task repository")
	}

	return GrpcServer{
		taskRepo: taskRepo,
	}
}

func MapTaskToProto(task app.Task) *taskv1.Task {
	if task.Description == nil {
		task.Description = new(string)
	}

	return &taskv1.Task{
		TaskId:      task.TaskId,
		Title:       task.Title,
		Description: *task.Description,
		CreatedAt:   timestamppb.New(task.CreatedAt.Time),
		UpdatedAt:   timestamppb.New(task.UpdatedAt.Time),
	}
}

func (g GrpcServer) CreateTask(ctx context.Context, req *taskv1.CreateTaskRequest) (*taskv1.CreateTaskResponse, error) {
	log := logger.Get()

	createTaskDto := &app.CreateTaskDto{
		Title:       req.Body.Title,
		Description: req.Body.Description,
	}

	if err := validate.Validate(createTaskDto); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	createdTask, err := g.taskRepo.CreateTask(ctx, createTaskDto)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return nil, status.Error(codes.AlreadyExists, "task already exists")
			case pgerrcode.NotNullViolation:
				return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("missing required field: %s", pgErr.ColumnName))
			}
		}

		log.Error().Err(err).Msg("failed to create task")

		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &taskv1.CreateTaskResponse{
		Task: MapTaskToProto(*createdTask),
	}, nil
}

func (g GrpcServer) GetAllTask(ctx context.Context, req *taskv1.GetAllTaskRequest) (*taskv1.GetAllTaskResponse, error) {
	log := logger.Get()

	tasks, err := g.taskRepo.GetAllTask(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get tasks")

		return nil, status.Error(codes.Internal, "failed to get tasks")
	}

	var taskList []*taskv1.Task
	for _, task := range tasks {
		taskList = append(taskList, MapTaskToProto(*task))
	}

	return &taskv1.GetAllTaskResponse{
		Tasks: taskList,
	}, nil
}

func (g GrpcServer) GetTaskById(ctx context.Context, req *taskv1.GetTaskByIdRequest) (*taskv1.GetTaskByIdResponse, error) {
	log := logger.Get()

	task, err := g.taskRepo.GetTaskById(ctx, req.TaskId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "task not found")
		}

		log.Error().Err(err).Msg("failed to get task")

		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &taskv1.GetTaskByIdResponse{
		Task: MapTaskToProto(*task),
	}, nil
}

func (g GrpcServer) UpdateTask(ctx context.Context, req *taskv1.UpdateTaskRequest) (*taskv1.UpdateTaskResponse, error) {
	log := logger.Get()

	updateTaskDto := &app.UpdateTaskDto{
		Title:       req.Body.Title,
		Description: req.Body.Description,
	}

	if err := validate.Validate(updateTaskDto); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	updatedTask, err := g.taskRepo.UpdateTask(ctx, req.TaskId, updateTaskDto)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "task not found")
		}

		log.Error().Err(err).Msg("failed to update task")

		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &taskv1.UpdateTaskResponse{
		Task: MapTaskToProto(*updatedTask),
	}, nil
}

func (g GrpcServer) DeleteTask(ctx context.Context, req *taskv1.DeleteTaskRequest) (*taskv1.DeleteTaskResponse, error) {
	log := logger.Get()

	deletedTask, err := g.taskRepo.DeleteTask(ctx, req.TaskId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "task not found")
		}

		log.Error().Err(err).Msg("failed to delete task")

		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &taskv1.DeleteTaskResponse{
		Task: MapTaskToProto(*deletedTask),
	}, nil
}
