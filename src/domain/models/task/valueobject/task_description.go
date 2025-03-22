package valueobject

import (
	"fmt"
	"unicode/utf8"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

const (
	TASK_DESCRIPTION_MIN_LENGTH = 1
	TASK_DESCRIPTION_MAX_LENGTH = 200
)

type TaskDescription struct {
	Core.ValueObject[string]
}

func NewTaskDescription(value string) (*TaskDescription, error) {
	if utf8.RuneCountInString(value) < TASK_DESCRIPTION_MIN_LENGTH {
		return nil, fmt.Errorf("TaskDescription must be at least %d characters long", TASK_DESCRIPTION_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > TASK_DESCRIPTION_MAX_LENGTH {
		return nil, fmt.Errorf("TaskDescription must be no more than %d characters long", TASK_DESCRIPTION_MAX_LENGTH)
	}
	return &TaskDescription{ValueObject: Core.NewValueObject(value)}, nil
}
