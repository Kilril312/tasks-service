package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/Kilril312/project-protos/proto/task"
	userpb "github.com/Kilril312/project-protos/proto/user"
	"github.com/Kilril312/tasks-service/internal/task"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	createdTask, err := h.svc.CreateTask(ctx, &task.Task{
		UserId: uint(req.UserId),
		Title:  req.Title,
	})

	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(createdTask.ID),
			UserId: uint32(createdTask.UserId),
			Title:  createdTask.Title,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	taskId, err := h.svc.GetTask(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(taskId.ID),
			UserId: uint32(taskId.UserId),
			Title:  taskId.Title,
		},
	}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	updTask := &task.Task{
		ID:    uint(req.Id),
		Title: req.NewTitle,
	}

	if err := h.svc.UpdateTask(ctx, updTask); err != nil {
		return nil, err
	}

	updatedTask, err := h.svc.GetTask(ctx, uint(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to get updatedtask: %w", err)
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(updatedTask.ID),
			UserId: uint32(updatedTask.UserId),
			Title:  updTask.Title,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTask(ctx, uint(req.Id)); err != nil {
		return nil, err
	}
	return &taskpb.DeleteTaskResponse{
		Success: true,
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, _ *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTask(ctx)
	if err != nil {
		return nil, err
	}

	list := &taskpb.ListTasksResponse{
		Tasks: make([]*taskpb.Task, 0, len(tasks)),
	}

	for _, u := range tasks {
		list.Tasks = append(list.Tasks, &taskpb.Task{
			Id:     uint32(u.ID),
			UserId: uint32(u.UserId),
			Title:  u.Title,
		})
	}

	return list, nil
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	tasks, err := h.svc.GetTaskbyUserID(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	list := &taskpb.ListTasksResponse{
		Tasks: make([]*taskpb.Task, 0, len(tasks)),
	}

	for _, u := range tasks {
		list.Tasks = append(list.Tasks, &taskpb.Task{
			Id:     uint32(u.ID),
			UserId: uint32(u.UserId),
			Title:  u.Title,
		})
	}

	return list, nil
}
