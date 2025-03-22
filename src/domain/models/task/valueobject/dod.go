package valueobject

import (
	"fmt"
	"unicode/utf8"

	Core "github.com/yoshi-d-24/goal-sync/domain/models/core"
)

const (
	DOD_MIN_LENGTH = 1
	DOD_MAX_LENGTH = 300
)

type DoD struct {
	Core.ValueObject[string]
}

func NewDoD(value string) (*DoD, error) {
	if utf8.RuneCountInString(value) < DOD_MIN_LENGTH {
		return nil, fmt.Errorf("DoD must be at least %d characters long", DOD_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > DOD_MAX_LENGTH {
		return nil, fmt.Errorf("DoD must be no more than %d characters long", DOD_MAX_LENGTH)
	}
	return &DoD{ValueObject: Core.NewValueObject(value)}, nil
}
