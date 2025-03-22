package value

import (
	"fmt"
	"unicode/utf8"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

const (
	TASK_COMMENT_MIN_LENGTH = 1
	TASK_COMMENT_MAX_LENGTH = 1000
)

type TaskComment struct {
	Core.IValueObject[string]
}

func NewTaskComment(value string) (*TaskComment, error) {
	if utf8.RuneCountInString(value) < DOD_MIN_LENGTH {
		return nil, fmt.Errorf("TaskComment must be at least %d characters long", DOD_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > DOD_MAX_LENGTH {
		return nil, fmt.Errorf("TaskComment must be no more than %d characters long", DOD_MAX_LENGTH)
	}
	return &TaskComment{Core.NewValueObject(value)}, nil
}
