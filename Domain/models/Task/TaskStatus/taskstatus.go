package task

import (
	"fmt"
)

type TaskStatus int

const (
	Incomplete TaskStatus = iota
	InProgress
	Complete
)

type TaskStatusInterface interface {
	Value() int
}

func NewTaskStatus(value int) (*TaskStatus, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	taskStatus := TaskStatus(value)
	return &taskStatus, nil
}

func (t *TaskStatus) Value() int {
	return int(*t)
}

func validate(value int) error {
	if value < int(Incomplete) || value > int(Complete) {
		return fmt.Errorf("invalid TaskStatus value: %d", value)
	}
	return nil
}
