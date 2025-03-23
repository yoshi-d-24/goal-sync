package task

import (
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
)

type InMemoryTaskRepository struct {
	TaskModel.ITaskRepository
}

var DB = map[string]*TaskModel.Task{}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{}
}

func (r *InMemoryTaskRepository) FindById(id string) (*TaskModel.Task, error) {
	if task, ok := DB[id]; ok {
		return task, nil
	}

	return nil, nil
}

func (r *InMemoryTaskRepository) FindByTitle(title string) (*TaskModel.Task, error) {
	for _, task := range DB {
		if task.Title().Value() == title {
			return task, nil
		}
	}

	return nil, nil
}

func (r *InMemoryTaskRepository) FindAll() ([]*TaskModel.Task, error) {
	tasks := []*TaskModel.Task{}
	for _, task := range DB {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) Save(task *TaskModel.Task) error {
	DB[task.Id().Value()] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	delete(DB, id)
	return nil
}
