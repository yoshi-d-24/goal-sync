package task

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) FindById(id string) (*TaskModel.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*TaskModel.Task), args.Error(1)
}

func (m *MockTaskRepository) FindByTitle(title string) (*TaskModel.Task, error) {
	args := m.Called(title)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*TaskModel.Task), args.Error(1)
}

func (m *MockTaskRepository) FindAll() ([]*TaskModel.Task, error) {
	args := m.Called()
	return args.Get(0).([]*TaskModel.Task), args.Error(1)
}

func (m *MockTaskRepository) Save(task *TaskModel.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockTaskDomainService struct {
	mock.Mock
}

func (m *MockTaskDomainService) TaskRepository() TaskModel.ITaskRepository {
	return nil
}

func (m *MockTaskDomainService) ExistsDuplicateTitle(task *TaskModel.Task) (bool, error) {
	args := m.Called(task)
	return args.Get(0).(bool), args.Error(1)
}

const (
	uuid_1 = "50ac2aa3-ab64-4184-9112-d23221dc1832"
)

var TestUUIDGenerator = func() uuid.UUID {
	id, _ := uuid.Parse(uuid_1)
	return id
}

func DefaultCommand() RegisterTaskCommand {
	return RegisterTaskCommand{
		Title:       "test title",
		Description: "test description",
		Dod:         "test dod",
	}
}

func DefaultExpectedTask() *TaskModel.Task {
	taskID, _ := VO.NewTaskId(uuid_1)
	title, _ := VO.NewTitle("test title")
	description, _ := VO.NewTaskDescription("test description")
	status, _ := VO.NewTaskStatus(VO.Incomplete)
	dod, _ := VO.NewDoD("test dod")

	return TaskModel.NewTask(*taskID, *title, *description, *status, *dod)
}

func TestRegisterTaskApplicationService_Execute(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockDomainService := new(MockTaskDomainService)

	type args struct {
		command RegisterTaskCommand
	}
	tests := []struct {
		name          string
		args          args
		expectedError error
		setup         func(a args)
	}{
		{
			name: "正常系: タスクの登録",
			args: args{
				command: DefaultCommand(),
			},
			expectedError: nil,
			setup: func(a args) {
				expectedTask := DefaultExpectedTask()

				mockRepo.On("Save", expectedTask).Return(nil)
				mockDomainService.On("ExistsDuplicateTitle", expectedTask).Return(false, nil)
			},
		},
		{
			name: "異常系: タイトル重複",
			args: args{
				command: DefaultCommand(),
			},
			expectedError: fmt.Errorf("task title is duplicated"),
			setup: func(a args) {
				expectedTask := DefaultExpectedTask()

				// タイトル重複をシミュレート
				mockRepo.On("Save", expectedTask).Return(nil)
				mockDomainService.On("ExistsDuplicateTitle", expectedTask).Return(true, nil)
			},
		},
		{
			name: "異常系: リポジトリのエラー",
			args: args{
				command: DefaultCommand(),
			},
			expectedError: fmt.Errorf("test error"),
			setup: func(a args) {
				expectedTask := DefaultExpectedTask()

				mockRepo.On("Save", expectedTask).Return(fmt.Errorf("test error"))
				mockDomainService.On("ExistsDuplicateTitle", expectedTask).Return(false, nil)
			},
		},
		{
			name: "異常系: ドメインサービスのエラー",
			args: args{
				command: DefaultCommand(),
			},
			expectedError: fmt.Errorf("test error"),
			setup: func(a args) {
				expectedTask := DefaultExpectedTask()

				mockRepo.On("Save", expectedTask).Return(nil)
				mockDomainService.On("ExistsDuplicateTitle", expectedTask).Return(false, fmt.Errorf("test error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.args)
			}
			s := RegisterTaskApplicationService{repository: mockRepo, domainService: mockDomainService, uuidGenerator: TestUUIDGenerator}
			if err := s.Execute(tt.args.command); err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("RegisterTaskApplicationService.execute() error = %v, wantErr %v", err, tt.expectedError)
			}
		})
	}
}
