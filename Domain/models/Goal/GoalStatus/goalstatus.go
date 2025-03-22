package goal

import (
	"fmt"
)

type GoalStatus int

const (
	Incomplete GoalStatus = iota
	Complete
)

type GoalStatusInterface interface {
	Value() int
}

func NewGoalStatus(value int) (GoalStatusInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}
	goalStatus := GoalStatus(value)
	return &goalStatus, nil
}

func (g *GoalStatus) Value() int {
	return int(*g)
}

func validate(value int) error {
	if value < int(Incomplete) || value > int(Complete) {
		return fmt.Errorf("invalid GoalStatusEnum value: %d", value)
	}

	return nil
}
