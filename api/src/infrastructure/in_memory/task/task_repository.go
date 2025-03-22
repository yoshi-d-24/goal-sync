package task

import (
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
)

type InMemoryTaskRepository struct {
	TaskModel.ITaskRepository
}

var DB = map[int]*TaskModel.Task{}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{}
}

func (r *InMemoryTaskRepository) FindById(id int) (*TaskModel.Task, error) {
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

func (r *InMemoryTaskRepository) Delete(task *TaskModel.Task) error {
	delete(DB, task.Id().Value())
	return nil
}
