package goal

import (
	"fmt"
	"unicode/utf8"
)

const (
	MIN_LENGTH = 1
	MAX_LENGTH = 500
)

type GoalDescriptionInterface interface {
	Value() string
}

type GoalDescription struct {
	value string
}

func NewGoalDescription(value string) (GoalDescriptionInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	return &GoalDescription{value}, nil
}

func (g *GoalDescription) Value() string {
	return g.value
}

func validate(value string) error {
	if utf8.RuneCountInString(value) < MIN_LENGTH {
		return fmt.Errorf("goal description must be at least %d characters long", MIN_LENGTH)
	}
	if utf8.RuneCountInString(value) > MAX_LENGTH {
		return fmt.Errorf("goal description must be no more than %d characters long", MAX_LENGTH)
	}
	return nil
}
