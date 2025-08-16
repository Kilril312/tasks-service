package task

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	err := s.repo.Create(task)

	return task, err
}

func (s *Service) GetTaskbyID(ctx context.Context, id uint) (*Task, error) {
	return s.repo.GetbyID(id)
}

func (s *Service) GetTaskbyUserID(ctx context.Context, user_id uint) ([]Task, error) {
	return s.repo.GetTaskbyUserID(user_id)
}

func (s *Service) UpdateTask(ctx context.Context, task *Task) error {
	return s.repo.Update(task)
}

func (s *Service) DeleteTask(ctx context.Context, id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) ListTask(ctx context.Context) ([]Task, error) {
	return s.repo.List()
}
