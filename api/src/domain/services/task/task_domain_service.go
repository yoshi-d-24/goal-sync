package task

import (
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
)

type TaskDomainService struct {
	taskRepository TaskModel.ITaskRepository
}

func NewTaskDomainService(repository TaskModel.ITaskRepository) *TaskDomainService {
	return &TaskDomainService{taskRepository: repository}
}

func (ts *TaskDomainService) ExistsDuplicateTitle(task *TaskModel.Task) (bool, error) {
	duplicateTitleTask, err := ts.taskRepository.FindByTitle((*task).Title().Value())
	if err != nil {
		return false, err
	}
	if duplicateTitleTask != nil {
		return false, err
	}
	return true, nil
}
