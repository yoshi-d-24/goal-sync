package valueobject

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
	if value < int(Incomplete) || value > int(Complete) {
		return nil, fmt.Errorf("invalid TaskStatus value: %d", value)
	}

	taskStatus := TaskStatus(value)
	return &taskStatus, nil
}

func (t *TaskStatus) Value() int {
	return int(*t)
}
