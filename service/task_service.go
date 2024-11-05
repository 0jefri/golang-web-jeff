package service

import (
	"fmt"

	"github.com/golang-web/model"
	"github.com/golang-web/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) RegisterNewTask(payload model.Task) error {
	if payload.ID == "" || payload.Name == "" || payload.Status == "" {
		return fmt.Errorf("all payload is required")
	}

	err := s.taskRepo.Create(&payload)
	if err != nil {
		return fmt.Errorf("failed to create task: %s", err)
	}
	return nil
}

func (s *TaskService) GetAllTasks() ([]*model.Task, error) {
	tasks, err := s.taskRepo.GetAllTask()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %v", err)
	}
	return tasks, nil
}
