package valueobject

import (
	"fmt"
	"unicode/utf8"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

const (
	TITLE_MIN_LENGTH = 1
	TITLE_MAX_LENGTH = 50
)

type Title struct {
	Core.ValueObject[string]
}

func NewTitle(value string) (*Title, error) {
	if utf8.RuneCountInString(value) < TITLE_MIN_LENGTH {
		return nil, fmt.Errorf("title must be at least %d characters long", TITLE_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > TITLE_MAX_LENGTH {
		return nil, fmt.Errorf("title must be no more than %d characters long", TITLE_MAX_LENGTH)
	}
	return &Title{ValueObject: Core.NewValueObject(value)}, nil
}
