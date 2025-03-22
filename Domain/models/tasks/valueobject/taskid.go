package valueobject

import (
	"fmt"
)

type TaskId struct {
	value int
}

func NewTaskId(value int) (*TaskId, error) {
	if value <= 0 {
		return nil, fmt.Errorf("TaskId must be greater than 0")
	}
	return &TaskId{value}, nil
}

func (t *TaskId) Value() int {
	return t.value
}
