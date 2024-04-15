package ports

import (
	"context"

	tasksv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/tasks/v1"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/app"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	taskRepository app.TaskRepository
	tasksv1.UnimplementedTaskServiceServer
}

func NewGrpcServer(taskRepository app.TaskRepository) GrpcServer {
	return GrpcServer{
		taskRepository: taskRepository,
	}
}

func (g GrpcServer) GetAllTasks(ctx context.Context, _ *tasksv1.GetAllTasksRequest) (*tasksv1.GetAllTasksResponse, error) {
	tasks, err := g.taskRepository.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := make([]*tasksv1.Task, 0)

	for _, task := range tasks {
		response = append(response, &tasksv1.Task{
			Id:          task.UUID,
			Title:       task.Title,
			Description: task.Description,
		})
	}

	return &tasksv1.GetAllTasksResponse{
		Tasks: response,
	}, nil
}

func (g GrpcServer) CreateTask(ctx context.Context, request *tasksv1.CreateTaskRequest) (*tasksv1.CreateTaskResponse, error) {
	err := g.taskRepository.Create(ctx, &app.Task{
		UUID:        uuid.New().String(),
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &tasksv1.CreateTaskResponse{}, nil
}

func (g GrpcServer) DeleteTask(ctx context.Context, request *tasksv1.DeleteTaskRequest) (*tasksv1.DeleteTaskResponse, error) {
	err := g.taskRepository.Delete(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &tasksv1.DeleteTaskResponse{}, nil
}

func (g GrpcServer) GetOneTask(ctx context.Context, request *tasksv1.GetOneTaskRequest) (*tasksv1.GetOneTaskResponse, error) {
	task, err := g.taskRepository.GetById(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &tasksv1.GetOneTaskResponse{
		Task: &tasksv1.Task{
			Id:          task.UUID,
			Title:       task.Title,
			Description: task.Description,
		},
	}, nil
}

func (g GrpcServer) UpdateTask(ctx context.Context, request *tasksv1.UpdateTaskRequest) (*tasksv1.UpdateTaskResponse, error) {
	err := g.taskRepository.Update(ctx, request.Id, &app.Task{
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &tasksv1.UpdateTaskResponse{}, nil
}
