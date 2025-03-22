package task

import (
	"testing"

	"github.com/stretchr/testify/assert"

	VO "github.com/yoshi-d-24/goal-sync/domain/models/task/value"
)

func TestNewTask(t *testing.T) {
	assert := assert.New(t)

	id, _ := VO.NewTaskId(1)
	title, _ := VO.NewTitle("Task Title")
	description, _ := VO.NewTaskDescription("Task Description")
	status, _ := VO.NewTaskStatus(VO.Incomplete)
	dod, _ := VO.NewDoD("DoD")

	task := NewTask(*id, *title, *description, *status, *dod)

	assert.NotNil(task)
	assert.True(id.Equals(task.Id()))
	assert.True(title.Equals(task.Title()))
	assert.True(description.Equals(task.Description()))
	assert.True(status.Equals(task.Status()))
	assert.True(dod.Equals(task.DoD()))
}
