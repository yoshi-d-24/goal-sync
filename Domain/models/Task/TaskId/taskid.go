package taskid

import (
	"fmt"
)

type TaskId struct {
	value int
}

func NewTaskId(value int) (*TaskId, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	return &TaskId{value}, nil
}

func (t *TaskId) Value() int {
	return t.value
}

func validate(value int) error {
	if value <= 0 {
		return fmt.Errorf("TaskId must be greater than 0")
	}
	return nil
}
