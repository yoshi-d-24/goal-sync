package valueobject

import (
	"fmt"
	"unicode/utf8"
)

const (
	TASK_DESCRIPTION_MIN_LENGTH = 1
	TASK_DESCRIPTION_MAX_LENGTH = 200
)

type TaskDescriptionInterface interface {
	Value() string
}

type TaskDescription struct {
	value string
}

func NewTaskDescription(value string) (TaskDescriptionInterface, error) {
	if utf8.RuneCountInString(value) < TASK_DESCRIPTION_MIN_LENGTH {
		return nil, fmt.Errorf("TaskDescription must be at least %d characters long", TASK_DESCRIPTION_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > TASK_DESCRIPTION_MAX_LENGTH {
		return nil, fmt.Errorf("TaskDescription must be no more than %d characters long", TASK_DESCRIPTION_MAX_LENGTH)
	}
	return &TaskDescription{value}, nil
}

func (t *TaskDescription) Value() string {
	return t.value
}

func validate(value string) error {

	return nil
}
