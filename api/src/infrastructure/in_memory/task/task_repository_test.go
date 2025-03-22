package task

import (
	"reflect"
	"testing"

	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

func TestInMemoryTaskRepository_FindById(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	id := 1
	task := createTestTaskWithTitle(id, "test task")
	repo.Save(task)

	type args struct {
		id int
	}
	tests := []struct {
		name        string
		args        args
		expected    *TaskModel.Task
		expectedErr bool
	}{
		{
			name:        "正常系",
			args:        args{id: id},
			expected:    task,
			expectedErr: false,
		},
		{
			name:        "存在しないID",
			args:        args{id: 2},
			expected:    nil,
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindById(tt.args.id)
			if (err != nil) != tt.expectedErr {
				t.Errorf("InMemoryTaskRepository.FindById() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				if tt.expected != nil && got != nil {
					if got.Id().Value() != tt.expected.Id().Value() {
						t.Errorf("InMemoryTaskRepository.FindById() = %v, expected %v", got.Id().Value(), tt.expected.Id().Value())
					}
				} else {
					t.Errorf("InMemoryTaskRepository.FindById() = %v, expected %v", got, tt.expected)
				}
			}
		})
	}

	repo.Delete(task)
}

func TestInMemoryTaskRepository_FindByTitle(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	id := 1
	title := "test task"
	task := createTestTaskWithTitle(id, title)
	repo.Save(task)

	type args struct {
		title string
	}
	tests := []struct {
		name        string
		args        args
		expected    *TaskModel.Task
		expectedErr bool
	}{
		{
			name:        "正常系",
			args:        args{title: title},
			expected:    task,
			expectedErr: false,
		},
		{
			name:        "存在しないタイトル",
			args:        args{title: "not found"},
			expected:    nil,
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindByTitle(tt.args.title)
			if (err != nil) != tt.expectedErr {
				t.Errorf("InMemoryTaskRepository.FindByTitle() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				if tt.expected != nil && got != nil {
					if got.Id().Value() != tt.expected.Id().Value() {
						t.Errorf("InMemoryTaskRepository.FindByTitle() = %v, expected %v", got.Id().Value(), tt.expected.Id().Value())
					}
				} else {
					t.Errorf("InMemoryTaskRepository.FindByTitle() = %v, expected %v", got, tt.expected)
				}
			}
		})
	}

	repo.Delete(task)
}

func TestInMemoryTaskRepository_FindAll(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	id := 1
	title := "test task"
	task := createTestTaskWithTitle(id, title)
	repo.Save(task)

	tests := []struct {
		name        string
		expected    []*TaskModel.Task
		expectedErr bool
	}{
		{
			name:        "正常系",
			expected:    []*TaskModel.Task{task},
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindAll()
			if (err != nil) != tt.expectedErr {
				t.Errorf("InMemoryTaskRepository.FindAll() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				if len(got) != len(tt.expected) {
					t.Errorf("InMemoryTaskRepository.FindAll() = %v, expected %v", len(got), len(tt.expected))
				} else {
					if got[0].Id().Value() != tt.expected[0].Id().Value() {
						t.Errorf("InMemoryTaskRepository.FindAll() = %v, expected %v", got[0].Id().Value(), tt.expected[0].Id().Value())
					}
				}
			}
		})
	}

	repo.Delete(task)
}

func TestInMemoryTaskRepository_Save(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	id := 1
	title := "test task"
	task := createTestTaskWithTitle(id, title)

	type args struct {
		task *TaskModel.Task
	}
	tests := []struct {
		name        string
		args        args
		expectedErr bool
	}{
		{
			name:        "正常系",
			args:        args{task: task},
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Save(tt.args.task); (err != nil) != tt.expectedErr {
				t.Errorf("InMemoryTaskRepository.Save() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			savedTask, _ := repo.FindById(id)
			if !reflect.DeepEqual(savedTask, task) {
				if task != nil && savedTask != nil {
					if savedTask.Id().Value() != task.Id().Value() {
						t.Errorf("InMemoryTaskRepository.Save() = %v, expected %v", savedTask.Id().Value(), task.Id().Value())
					}
				} else {
					t.Errorf("InMemoryTaskRepository.Save() = %v, expected %v", savedTask, task)
				}
			}
		})
	}
}

func TestInMemoryTaskRepository_Delete(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	id := 1
	title := "test task"
	task := createTestTaskWithTitle(id, title)
	repo.Save(task)

	type args struct {
		task *TaskModel.Task
	}
	tests := []struct {
		name        string
		args        args
		expectedErr bool
	}{
		{
			name:        "正常系",
			args:        args{task: task},
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Delete(tt.args.task); (err != nil) != tt.expectedErr {
				t.Errorf("InMemoryTaskRepository.Delete() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			deletedTask, _ := repo.FindById(id)
			if deletedTask != nil {
				t.Errorf("InMemoryTaskRepository.Delete() task not deleted")
			}
		})
	}
}

func createTestTaskWithTitle(id int, title string) *TaskModel.Task {
	taskId, _ := VO.NewTaskId(id)
	titleValue, _ := VO.NewTitle(title)
	description, _ := VO.NewTaskDescription("Task Description")
	status, _ := VO.NewTaskStatus(VO.Incomplete)
	dod, _ := VO.NewDoD("DoD")

	return TaskModel.NewTask(
		*taskId,
		*titleValue,
		*description,
		*status,
		*dod,
	)
}
