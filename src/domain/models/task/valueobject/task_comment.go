package valueobject

import (
	"fmt"
	"unicode/utf8"
)

const (
	TASK_COMMENT_MIN_LENGTH = 1
	TASK_COMMENT_MAX_LENGTH = 1000
)

type TaskCommentInterface interface {
	Value() string
}

type TaskComment struct {
	value string
}

func NewTaskComment(value string) (TaskCommentInterface, error) {
	if utf8.RuneCountInString(value) < DOD_MIN_LENGTH {
		return nil, fmt.Errorf("TaskComment must be at least %d characters long", DOD_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > DOD_MAX_LENGTH {
		return nil, fmt.Errorf("TaskComment must be no more than %d characters long", DOD_MAX_LENGTH)
	}
	return &TaskComment{value}, nil
}

func (t *TaskComment) Value() string {
	return t.value
}
