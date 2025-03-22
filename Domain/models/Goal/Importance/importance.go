package goal

import (
	"fmt"
)

const MIN = 1
const MAX = 10

type ImportanceInterface interface {
	Value() int
}

type Importance struct {
	value int
}

func NewGoalStatus(value int) (*Importance, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	return &Importance{value}, nil
}

func (i *Importance) Value() int {
	return i.value
}

func validate(value int) error {
	if value < MIN || value > MAX {
		return fmt.Errorf("importance must be between %d and %d", MIN, MAX)
	}

	return nil
}
