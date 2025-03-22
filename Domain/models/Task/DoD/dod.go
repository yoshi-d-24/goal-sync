package task

import (
	"fmt"
	"unicode/utf8"
)

const (
	MIN_LENGTH = 1
	MAX_LENGTH = 300
)

type DoDInterface interface {
	Value() string
}

type DoD struct {
	value string
}

func NewDoD(value string) (DoDInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	return &DoD{value}, nil
}

func (d *DoD) Value() string {
	return d.value
}

func validate(value string) error {
	if utf8.RuneCountInString(value) < MIN_LENGTH {
		return fmt.Errorf("DoD must be at least %d characters long", MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > MAX_LENGTH {
		return fmt.Errorf("DoD must be no more than %d characters long", MAX_LENGTH)
	}
	return nil
}
