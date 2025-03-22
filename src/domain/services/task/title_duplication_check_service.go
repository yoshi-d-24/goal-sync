package task

import (
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
)

type TitleDuplicationCheckService struct {
	taskRepository TaskModel.ITaskRepository
}

func (ts *TitleDuplicationCheckService) ExistsDuplicateTitle(task *TaskModel.TaskInterface) (bool, error) {
	sameTitleTask, err := ts.taskRepository.FindByTitle((*task).Title().Value())
	if err != nil {
		return false, err
	}
	if sameTitleTask != nil {
		return false, err
	}
	return true, nil
}
