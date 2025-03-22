package goal

import (
	"fmt"
	"unicode/utf8"
)

const (
	MIN_LENGTH = 1
	MAX_LENGTH = 50
)

type Title struct {
	value string
}

func NewTitle(value string) (*Title, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	return &Title{value}, nil
}

func (t *Title) Value() string {
	return t.value
}

func validate(value string) error {
	if utf8.RuneCountInString(value) < MIN_LENGTH {
		return fmt.Errorf("title must be at least %d characters long", MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > MAX_LENGTH {
		return fmt.Errorf("title must be no more than %d characters long", MAX_LENGTH)
	}
	return nil
}
