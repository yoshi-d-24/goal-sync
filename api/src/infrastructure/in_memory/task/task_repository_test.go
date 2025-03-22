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
		name    string
		args    args
		want    *TaskModel.Task
		wantErr bool
	}{
		{
			name:    "正常系",
			args:    args{id: id},
			want:    task,
			wantErr: false,
		},
		{
			name:    "存在しないID",
			args:    args{id: 2},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryTaskRepository.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				if tt.want != nil && got != nil {
					if got.Id().Value() != tt.want.Id().Value() {
						t.Errorf("InMemoryTaskRepository.FindById() = %v, want %v", got.Id().Value(), tt.want.Id().Value())
					}
				} else {
					t.Errorf("InMemoryTaskRepository.FindById() = %v, want %v", got, tt.want)
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
		name    string
		args    args
		want    *TaskModel.Task
		wantErr bool
	}{
		{
			name:    "正常系",
			args:    args{title: title},
			want:    task,
			wantErr: false,
		},
		{
			name:    "存在しないタイトル",
			args:    args{title: "not found"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindByTitle(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryTaskRepository.FindByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				if tt.want != nil && got != nil {
					if got.Id().Value() != tt.want.Id().Value() {
						t.Errorf("InMemoryTaskRepository.FindByTitle() = %v, want %v", got.Id().Value(), tt.want.Id().Value())
					}
				} else {
					t.Errorf("InMemoryTaskRepository.FindByTitle() = %v, want %v", got, tt.want)
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
		name    string
		want    []*TaskModel.Task
		wantErr bool
	}{
		{
			name:    "正常系",
			want:    []*TaskModel.Task{task},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryTaskRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				if len(got) != len(tt.want) {
					t.Errorf("InMemoryTaskRepository.FindAll() = %v, want %v", len(got), len(tt.want))
				} else {
					if got[0].Id().Value() != tt.want[0].Id().Value() {
						t.Errorf("InMemoryTaskRepository.FindAll() = %v, want %v", got[0].Id().Value(), tt.want[0].Id().Value())
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
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "正常系",
			args:    args{task: task},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Save(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryTaskRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			savedTask, _ := repo.FindById(id)
			if !reflect.DeepEqual(savedTask, task) {
				if task != nil && savedTask != nil {
					if savedTask.Id().Value() != task.Id().Value() {
						t.Errorf("InMemoryTaskRepository.Save() = %v, want %v", savedTask.Id().Value(), task.Id().Value())
					}
				} else {
					t.Errorf("InMemoryTaskRepository.Save() = %v, want %v", savedTask, task)
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
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "正常系",
			args:    args{task: task},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Delete(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryTaskRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
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
