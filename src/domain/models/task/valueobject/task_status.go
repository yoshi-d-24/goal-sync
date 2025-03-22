package valueobject

import (
	"fmt"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

const (
	Incomplete int = iota
	InProgress
	Complete
)

type TaskStatus struct {
	Core.ValueObject[int]
}

func NewTaskStatus(value int) (*TaskStatus, error) {
	if value < Incomplete || value > Complete {
		return nil, fmt.Errorf("invalid TaskStatus value: %d", value)
	}

	taskStatus := TaskStatus{ValueObject: Core.NewValueObject(value)}
	return &taskStatus, nil
}
