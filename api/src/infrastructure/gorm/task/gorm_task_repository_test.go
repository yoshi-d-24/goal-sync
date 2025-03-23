package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	TaskModel "github.com/yoshi-d-24/goal-sync/domain/models/task"
	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
	GormCore "github.com/yoshi-d-24/goal-sync/infrastructure/gorm/core"
	GormTask "github.com/yoshi-d-24/goal-sync/infrastructure/gorm/task"
)

func setupTest(t *testing.T) (*gorm.DB, *GormTask.GormTaskRepository) {
	db := GormCore.CreateDB()
	db.Migrator().DropTable(&GormTask.TaskDto{})
	db.AutoMigrate(&GormTask.TaskDto{})
	repo := GormTask.NewGormTaskRepository(GormCore.CreateDB())
	return db, repo
}

func TestGormTaskRepository_FindById(t *testing.T) {
	assert := assert.New(t)
	db, repo := setupTest(t)

	// Create a task
	newTask := createTestTaskWithTitle(1, "test title")
	err := repo.Save(newTask)
	assert.NoError(err)

	// Find the task by ID
	foundTask, err := repo.FindById(1)
	assert.NoError(err)
	assert.NotNil(*foundTask)
	assert.True(newTask.Equals(foundTask))

	deleteTestData(db, []int{1})
}

func TestGormTaskRepository_FindAll(t *testing.T) {
	assert := assert.New(t)
	db, repo := setupTest(t)

	// Create tasks
	task1 := createTestTaskWithTitle(1, "test title 1")
	err := repo.Save(task1)
	assert.NoError(err)

	task2 := createTestTaskWithTitle(2, "test title 2")
	err = repo.Save(task2)
	assert.NoError(err)

	// Find all tasks
	foundTasks, err := repo.FindAll()
	assert.NoError(err)
	assert.Len(foundTasks, 2)
	assert.True(task1.Equals(foundTasks[0]))
	assert.True(task2.Equals(foundTasks[1]))

	deleteTestData(db, []int{1, 2})
}

func TestGormTaskRepository_FindByTitle(t *testing.T) {
	assert := assert.New(t)
	db, repo := setupTest(t)

	// Create a task
	title := "test title"
	newTask := createTestTaskWithTitle(1, title)
	err := repo.Save(newTask)
	assert.NoError(err)

	// Find the task by title
	foundTask, err := repo.FindByTitle(title)
	assert.NoError(err)
	assert.NotNil(foundTask)
	assert.True(newTask.Equals(foundTask))

	deleteTestData(db, []int{1})
}

func TestGormTaskRepository_Save(t *testing.T) {
	assert := assert.New(t)
	db, repo := setupTest(t)

	// Create a task
	newTask := createTestTaskWithTitle(1, "test title")

	// Save the task
	err := repo.Save(newTask)
	assert.NoError(err)

	// Verify the task is saved
	var savedTask GormTask.TaskDto
	db.First(&savedTask, 1)
	assert.Equal(newTask.Title().Value(), savedTask.Title)

	deleteTestData(db, []int{1})
}

func TestGormTaskRepository_Delete(t *testing.T) {
	assert := assert.New(t)
	db, repo := setupTest(t)

	// Create a task
	newTask := createTestTaskWithTitle(1, "test title")
	err := repo.Save(newTask)
	assert.NoError(err)

	// Delete the task
	err = repo.Delete(1)
	assert.NoError(err)

	// Verify the task is deleted
	foundTask, _ := repo.FindById(1)
	assert.Nil(foundTask)

	deleteTestData(db, []int{1})
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

func deleteTestData(db *gorm.DB, ids []int) {
	for _, id := range ids {
		db.Delete(&GormTask.TaskDto{}, id)
	}
}
