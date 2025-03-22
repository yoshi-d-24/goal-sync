package valueobject

import (
	"fmt"
	"unicode/utf8"
)

const (
	DOD_MIN_LENGTH = 1
	DOD_MAX_LENGTH = 300
)

type DoDInterface interface {
	Value() string
}

type DoD struct {
	value string
}

func NewDoD(value string) (DoDInterface, error) {
	if utf8.RuneCountInString(value) < DOD_MIN_LENGTH {
		return nil, fmt.Errorf("DoD must be at least %d characters long", DOD_MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > DOD_MAX_LENGTH {
		return nil, fmt.Errorf("DoD must be no more than %d characters long", DOD_MAX_LENGTH)
	}
	return &DoD{value}, nil
}

func (d *DoD) Value() string {
	return d.value
}
