package value

import (
	"fmt"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

type TaskId struct {
	Core.IValueObject[int]
}

func NewTaskId(value int) (*TaskId, error) {
	if value <= 0 {
		return nil, fmt.Errorf("TaskId must be greater than 0")
	}
	return &TaskId{Core.NewValueObject(value)}, nil
}
