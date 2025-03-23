package value

import (
	"fmt"

	"github.com/google/uuid"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

type TaskId struct {
	Core.IValueObject[string]
}

func NewTaskId(value string) (*TaskId, error) {
	err := uuid.Validate(value)

	if err != nil {
		return nil, fmt.Errorf("TaskId must be uuid")
	}
	return &TaskId{Core.NewValueObject(value)}, nil
}
