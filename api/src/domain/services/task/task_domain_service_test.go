package task

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) FindById(id int) (*TaskModel.Task, error) {
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

func (m *MockTaskRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTaskDomainService_ExistsDuplicateTitle(t *testing.T) {
	assert := assert.New(t)

	// テストケース1: 重複タイトルが存在する場合
	t.Run("duplicate title exists", func(t *testing.T) {
		mockRepo := new(MockTaskRepository)
		service := TaskDomainService{taskRepository: mockRepo}

		titleValue := "Duplicate Title"
		existingTask := createTestTaskWithTitle(1, titleValue)

		mockRepo.On("FindByTitle", titleValue).Return(existingTask, nil)

		task := createTestTaskWithTitle(2, titleValue)

		exists, err := service.ExistsDuplicateTitle(task)

		assert.False(exists)
		assert.NoError(err)
		mockRepo.AssertExpectations(t)
	})

	// テストケース2: 重複タイトルが存在しない場合
	t.Run("duplicate title does not exist", func(t *testing.T) {
		mockRepo := new(MockTaskRepository)
		service := TaskDomainService{taskRepository: mockRepo}

		titleValue := "New Title"
		mockRepo.On("FindByTitle", titleValue).Return(nil, nil)

		task := createTestTaskWithTitle(1, "New Title")

		exists, err := service.ExistsDuplicateTitle(task)

		assert.True(exists)
		assert.NoError(err)
		mockRepo.AssertExpectations(t)
	})

	// テストケース3: エラーが発生した場合
	t.Run("error occurs", func(t *testing.T) {
		mockRepo := new(MockTaskRepository)
		service := TaskDomainService{taskRepository: mockRepo}

		titleValue := "Some Title"
		expectedError := errors.New("repository error")

		mockRepo.On("FindByTitle", titleValue).Return(nil, expectedError)

		task := createTestTaskWithTitle(1, titleValue)

		exists, err := service.ExistsDuplicateTitle(task)

		assert.False(exists)
		assert.EqualError(err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})
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
