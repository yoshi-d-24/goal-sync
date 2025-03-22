package task

import (
	"fmt"
	"unicode/utf8"
)

const (
	MIN_LENGTH = 1
	MAX_LENGTH = 1000
)

type TaskCommentInterface interface {
	Value() string
}

type TaskComment struct {
	value string
}

func NewTaskComment(value string) (TaskCommentInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	return &TaskComment{value}, nil
}

func (t *TaskComment) Value() string {
	return t.value
}

func validate(value string) error {
	if utf8.RuneCountInString(value) < MIN_LENGTH {
		return fmt.Errorf("TaskComment must be at least %d characters long", MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > MAX_LENGTH {
		return fmt.Errorf("TaskComment must be no more than %d characters long", MAX_LENGTH)
	}
	return nil
}
