package goal

import (
	"fmt"
)

type GoalIdInterface interface {
	Value() int
}

type GoalId struct {
	value int
}

func NewGoalId(value int) (GoalIdInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}

	return &GoalId{value}, nil
}

func (g *GoalId) Value() int {
	return g.value
}

func validate(value int) error {
	if value < 1 {
		return fmt.Errorf("goal id should be greater than 0")
	}

	return nil
}
