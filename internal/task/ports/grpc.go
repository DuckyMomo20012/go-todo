package ports

import (
	"context"

	taskv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/task/v1"
	"github.com/DuckyMomo20012/go-todo/internal/task/app"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	taskRepository app.TaskRepository
	taskv1.UnimplementedTaskServiceServer
}

func NewGrpcServer(taskRepository app.TaskRepository) GrpcServer {
	return GrpcServer{
		taskRepository: taskRepository,
	}
}

func (g GrpcServer) GetAllTasks(ctx context.Context, _ *taskv1.GetAllTasksRequest) (*taskv1.GetAllTasksResponse, error) {
	tasks, err := g.taskRepository.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := make([]*taskv1.Task, 0)

	for _, task := range tasks {
		response = append(response, &taskv1.Task{
			Id:          task.UUID,
			Title:       task.Title,
			Description: task.Description,
		})
	}

	return &taskv1.GetAllTasksResponse{
		Tasks: response,
	}, nil
}

func (g GrpcServer) CreateTask(ctx context.Context, request *taskv1.CreateTaskRequest) (*taskv1.CreateTaskResponse, error) {
	err := g.taskRepository.Create(ctx, &app.Task{
		UUID:        uuid.New().String(),
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &taskv1.CreateTaskResponse{}, nil
}

func (g GrpcServer) DeleteTask(ctx context.Context, request *taskv1.DeleteTaskRequest) (*taskv1.DeleteTaskResponse, error) {
	err := g.taskRepository.Delete(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &taskv1.DeleteTaskResponse{}, nil
}

func (g GrpcServer) GetOneTask(ctx context.Context, request *taskv1.GetOneTaskRequest) (*taskv1.GetOneTaskResponse, error) {
	task, err := g.taskRepository.GetByID(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &taskv1.GetOneTaskResponse{
		Task: &taskv1.Task{
			Id:          task.UUID,
			Title:       task.Title,
			Description: task.Description,
		},
	}, nil
}

func (g GrpcServer) UpdateTask(ctx context.Context, request *taskv1.UpdateTaskRequest) (*taskv1.UpdateTaskResponse, error) {
	err := g.taskRepository.Update(ctx, request.Id, &app.Task{
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &taskv1.UpdateTaskResponse{}, nil
}
